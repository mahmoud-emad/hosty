package cmd

import (
	"strconv"

	"github.com/mahmoud-emad/hosty/internal/host"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved hosts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listHosts()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listHosts() error {
	hosty, err := host.NewHosty()
	if err != nil {
		return err
	}

	hosts, err := hosty.List()
	if err != nil {
		return err
	}

	printHeader()

	for _, h := range hosts {
		user := "-"
		if h.User != "" {
			user = h.User
		}

		port := "-"
		if h.Port > 0 {
			port = strconv.Itoa(h.Port)
		}

		printHost(h.Name, h.Address, port, user)
	}

	return nil
}
