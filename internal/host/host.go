package host

import (
	"fmt"
	"os"
	"path/filepath"

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
