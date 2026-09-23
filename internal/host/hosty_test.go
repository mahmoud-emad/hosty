package host

import (
	"path/filepath"
	"testing"

	"github.com/mahmoud-emad/hosty/internal/config"
)

func TestNewHost(t *testing.T) {
	h, err := NewHost("router", "192.168.1.1", "root", 22)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if h.Name != "router" {
		t.Errorf("expected name router, got %s", h.Name)
	}

	if h.Address != "192.168.1.1" {
		t.Errorf("expected address 192.168.1.1, got %s", h.Address)
	}

	if h.User != "root" {
		t.Errorf("expected user root, got %s", h.User)
	}

	if h.Port != 22 {
		t.Errorf("expected port 22, got %d", h.Port)
	}
}

func TestNewHostInvalidIP(t *testing.T) {
	_, err := NewHost("router", "not-an-ip", "root", 22)

	if err == nil {
		t.Fatal("expected an error for invalid IP")
	}
}

func TestNewHostInvalidPort(t *testing.T) {
	_, err := NewHost("router", "192.168.1.1", "root", 70000)

	if err == nil {
		t.Fatal("expected an error for invalid port")
	}
}

func TestNewHostNegativePort(t *testing.T) {
	_, err := NewHost("router", "192.168.1.1", "root", -1)

	if err == nil {
		t.Fatal("expected an error for negative port")
	}
}

func TestHostySet(t *testing.T) {
	dir := t.TempDir()

	configFile := filepath.Join(dir, "hosts.json")
	cfg := config.NewConfig(configFile)

	hosty := &Hosty{
		config: cfg,
		hosts:  make(map[string]*Host),
	}

	h, err := NewHost("router", "192.168.1.1", "root", 22)
	if err != nil {
		t.Fatal(err)
	}

	err = hosty.Set(h)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestHostySetDuplicate(t *testing.T) {
	dir := t.TempDir()

	configFile := filepath.Join(dir, "hosts.json")
	cfg := config.NewConfig(configFile)

	hosty := &Hosty{
		config: cfg,
		hosts:  make(map[string]*Host),
	}

	h, err := NewHost("router", "192.168.1.1", "root", 22)
	if err != nil {
		t.Fatal(err)
	}

	if err := hosty.Set(h); err != nil {
		t.Fatal(err)
	}

	err = hosty.Set(h)

	if err == nil {
		t.Fatal("expected duplicate host error")
	}
}

func TestHostyGet(t *testing.T) {
	dir := t.TempDir()

	configFile := filepath.Join(dir, "hosts.json")
	cfg := config.NewConfig(configFile)

	hosty := &Hosty{
		config: cfg,
		hosts:  make(map[string]*Host),
	}

	original, err := NewHost("router", "192.168.1.1", "root", 22)
	if err != nil {
		t.Fatal(err)
	}

	if err := hosty.Set(original); err != nil {
		t.Fatal(err)
	}

	found, err := hosty.Get("router")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if found.Name != "router" {
		t.Errorf("expected router, got %s", found.Name)
	}

	if found.Address != "192.168.1.1" {
		t.Errorf("expected 192.168.1.1, got %s", found.Address)
	}
}

func TestHostyForget(t *testing.T) {
	dir := t.TempDir()

	configFile := filepath.Join(dir, "hosts.json")
	cfg := config.NewConfig(configFile)

	hosty := &Hosty{
		config: cfg,
		hosts:  make(map[string]*Host),
	}

	h, err := NewHost("router", "192.168.1.1", "root", 22)
	if err != nil {
		t.Fatal(err)
	}

	if err := hosty.Set(h); err != nil {
		t.Fatal(err)
	}

	if err := hosty.Forget("router"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = hosty.Get("router")

	if err == nil {
		t.Fatal("expected host to no longer exist")
	}
}

func TestHostyUpdateAddress(t *testing.T) {
	dir := t.TempDir()

	configFile := filepath.Join(dir, "hosts.json")
	cfg := config.NewConfig(configFile)

	hosty := &Hosty{
		config: cfg,
		hosts:  make(map[string]*Host),
	}

	h, err := NewHost("router", "192.168.1.1", "root", 22)
	if err != nil {
		t.Fatal(err)
	}

	if err := hosty.Set(h); err != nil {
		t.Fatal(err)
	}

	if _, err := hosty.Update(
		"router",
		"",
		"192.168.1.50",
		"",
		0,
	); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	updated, err := hosty.Get("router")
	if err != nil {
		t.Fatal(err)
	}

	if updated.Address != "192.168.1.50" {
		t.Errorf(
			"expected address 192.168.1.50, got %s",
			updated.Address,
		)
	}
}

func TestHostyUpdateName(t *testing.T) {
	dir := t.TempDir()

	configFile := filepath.Join(dir, "hosts.json")
	cfg := config.NewConfig(configFile)

	hosty := &Hosty{
		config: cfg,
		hosts:  make(map[string]*Host),
	}

	h, err := NewHost("router", "192.168.1.1", "root", 22)
	if err != nil {
		t.Fatal(err)
	}

	if err := hosty.Set(h); err != nil {
		t.Fatal(err)
	}

	_, err = hosty.Update(
		"router",
		"router2",
		"",
		"",
		0,
	)
	if err != nil {
		t.Fatal(err)
	}

	_, err = hosty.Get("router")
	if err == nil {
		t.Fatal("expected old name to not exist")
	}

	updated, err := hosty.Get("router2")
	if err != nil {
		t.Fatalf("expected new name to exist: %v", err)
	}

	if updated.Name != "router2" {
		t.Errorf("expected router2, got %s", updated.Name)
	}
}

func TestNewHostInvalid(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		port int
	}{
		{
			name: "invalid ip",
			ip:   "hello",
			port: 22,
		},
		{
			name: "negative port",
			ip:   "192.168.1.1",
			port: -1,
		},
		{
			name: "port too large",
			ip:   "192.168.1.1",
			port: 70000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewHost("router", tt.ip, "root", tt.port)

			if err == nil {
				t.Fatal("expected an error")
			}
		})
	}
}
