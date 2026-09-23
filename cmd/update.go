package cmd

import (
	"strconv"

	"github.com/mahmoud-emad/hosty/internal/host"
	"github.com/spf13/cobra"
)

var (
	newName    string
	newAddress string
	newUser    string
	newPort    int
)

var updateCmd = &cobra.Command{
	Use:   "update <name> <--field>",
	Short: "Update host info",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateHost(
			args[0],
			newName,
			newAddress,
			newUser,
			newPort,
		)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&newName, "name", "n", "", "New host name")
	updateCmd.Flags().StringVarP(&newAddress, "address", "a", "", "New host address")
	updateCmd.Flags().StringVarP(&newUser, "user", "u", "", "New SSH username")
	updateCmd.Flags().IntVarP(&newPort, "port", "p", 0, "New SSH port")
}

func updateHost(
	oldName string,
	newName string,
	newAddress string,
	newUser string,
	newPort int,
) error {
	hosty, err := host.NewHosty()
	if err != nil {
		return err
	}

	h, err := hosty.Update(
		oldName,
		newName,
		newAddress,
		newUser,
		newPort,
	)
	if err != nil {
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

	printHeader()
	printHost(h.Name, h.Address, displayPort, displayUser)

	return nil
}
