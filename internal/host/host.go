package host

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/mahmoud-emad/hosty/internal/config"
)

const ConfigFilePath = ".config/hosty/hosts.json"

type Hosty struct {
	config *config.Config
	hosts  map[string]*Host
}

func NewHosty() (*Hosty, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configFile := filepath.Join(homeDir, ConfigFilePath)

	cfg := config.NewConfig(configFile)

	return &Hosty{
		config: cfg,
		hosts:  make(map[string]*Host),
	}, nil
}

func (h *Hosty) Set(host *Host) error {
	if err := h.config.Load(&h.hosts); err != nil {
		return err
	}

	if _, exists := h.hosts[host.Name]; exists {
		return fmt.Errorf("host %q already exists", host.Name)
	}

	h.hosts[host.Name] = host

	if err := h.config.Save(h.hosts); err != nil {
		return err
	}

	return nil
}

func (h *Hosty) Get(Name string) (*Host, error) {
	if err := h.config.Load(&h.hosts); err != nil {
		return nil, err
	}

	if _, exists := h.hosts[Name]; exists {
		return h.hosts[Name], nil
	}

	return nil, fmt.Errorf("host %q does not exists", Name)
}

func (h *Hosty) List() ([]*Host, error) {
	if err := h.config.Load(&h.hosts); err != nil {
		return nil, err
	}

	hosts := make([]*Host, 0, len(h.hosts))

	for _, host := range h.hosts {
		hosts = append(hosts, host)
	}

	return hosts, nil
}

func (h *Hosty) Forget(name string) error {
	if err := h.config.Load(&h.hosts); err != nil {
		return err
	}

	if _, exists := h.hosts[name]; !exists {
		return fmt.Errorf("host %q does not exist", name)
	}

	delete(h.hosts, name)

	if err := h.config.Save(h.hosts); err != nil {
		return err
	}

	return nil
}

func (h *Hosty) Update(oldName, newName, newAddress, newUser string, newPort int) (*Host, error) {
	if err := h.config.Load(&h.hosts); err != nil {
		return nil, err
	}

	host, exists := h.hosts[oldName]
	if !exists {
		return nil, fmt.Errorf("host %q does not exist", oldName)
	}

	if newName != "" {
		delete(h.hosts, oldName)

		host.Name = newName
		h.hosts[newName] = host
	}

	if newAddress != "" {
		host.Address = newAddress
	}

	if newUser != "" {
		host.User = newUser
	}

	if newPort != 0 {
		host.Port = newPort
	}

	if err := h.config.Save(h.hosts); err != nil {
		return nil, err
	}

	return host, nil
}

func (h *Hosty) Connect(name string) error {
	if err := h.config.Load(&h.hosts); err != nil {
		return err
	}

	host, exists := h.hosts[name]
	if !exists {
		return fmt.Errorf("host %q does not exist", name)
	}

	args := []string{}

	if host.Port > 0 {
		args = append(args, "-p", strconv.Itoa(host.Port))
	}

	target := host.Address

	if host.User != "" {
		target = fmt.Sprintf("%s@%s", host.User, host.Address)
	}

	args = append(args, target)
	fmt.Println("Connecting host", host.Address)

	cmd := exec.Command("ssh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
