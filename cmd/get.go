package cmd

import (
	"fmt"
	"strconv"

	"github.com/mahmoud-emad/hosty/internal/host"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <name>",
	Short: "Get a host by name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return getHost(args[0])
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getHost(name string) error {
	hosty, err := host.NewHosty()
	if err != nil {
		return err
	}

	h, err := hosty.Get(name)
	if err != nil {
		return err
	}

	user := "-"
	if len(h.User) > 0 {
		user = h.User
	}

	port := "-"
	if h.Port > 0 {
		port = strconv.Itoa(h.Port)
	}

	fmt.Printf("Name:    %s\n", h.Name)
	fmt.Printf("Address: %s\n", h.Address)
	fmt.Printf("Port:    %s\n", port)
	fmt.Printf("User:    %s\n", user)

	return nil
}
