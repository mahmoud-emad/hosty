package host

import (
	"fmt"
	"net"
)

type Host struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	User    string `json:"user,omitempty"`
	Port    int    `json:"port,omitempty"`
}

func NewHost(name, addr, user string, port int) (*Host, error) {
	if net.ParseIP(addr) == nil {
		return nil, fmt.Errorf("invalid IP address: %s", addr)
	}

	if port < 0 || port > 65535 {
		return nil, fmt.Errorf("invalid port: %d", port)
	}

	return &Host{
		Name:    name,
		Address: addr,
		User:    user,
		Port:    port,
	}, nil
}
