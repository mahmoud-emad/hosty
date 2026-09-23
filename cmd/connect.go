package cmd

import (
	"github.com/mahmoud-emad/hosty/internal/host"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect <name>",
	Short: "Connect to a host by name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return connectHost(args[0])
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}

func connectHost(name string) error {
	hosty, err := host.NewHosty()
	if err != nil {
		return err
	}

	return hosty.Connect(name)
}
