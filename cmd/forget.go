package cmd

import (
	"fmt"

	"github.com/mahmoud-emad/hosty/internal/host"
	"github.com/spf13/cobra"
)

var forgetCmd = &cobra.Command{
	Use:   "forget <name>",
	Short: "Delete a saved host by name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return forgetHost(args[0])
	},
}

func init() {
	rootCmd.AddCommand(forgetCmd)
}

func forgetHost(name string) error {
	hosty, err := host.NewHosty()
	if err != nil {
		return err
	}

	if err := hosty.Forget(name); err != nil {
		return err
	}

	fmt.Printf("Host %s deleted\n", name)
	return nil
}
