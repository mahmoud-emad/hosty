package host

import "testing"

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
