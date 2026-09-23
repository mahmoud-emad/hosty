package cmd

import (
	"fmt"
	"strconv"

	"github.com/mahmoud-emad/hosty/internal/host"
	"github.com/spf13/cobra"
)

var (
	user string
	port int
)

var setCmd = &cobra.Command{
	Use:   "set <name> <ip>",
	Short: "Add a host with a name and IP address",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return setHost(args[0], args[1], user, port)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringVarP(&user, "user", "u", "", "SSH username")
	setCmd.Flags().IntVarP(&port, "port", "p", 0, "SSH port")
}

func setHost(name, ip, user string, port int) error {
	hosty, err := host.NewHosty()
	if err != nil {
		return err
	}

	h, err := host.NewHost(name, ip, user, port)
	if err != nil {
		return err
	}

	if err := hosty.Set(h); err != nil {
		return err
	}

	displayUser := "-"
	if h.User != "" {
		displayUser = h.User
	}

	displayPort := "-"
	if h.Port > 0 {
		displayPort = strconv.Itoa(h.Port)
	}

	fmt.Printf("Name:    %s\n", h.Name)
	fmt.Printf("Address: %s\n", h.Address)
	fmt.Printf("Port:    %s\n", displayPort)
	fmt.Printf("User:    %s\n", displayUser)

	return nil
}
