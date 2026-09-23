// Package cmd wires up the Hosty command line.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hosty",
	Short: "Manage and connect to named hosts.",
	Long:  "Hosty is a small CLI tool for managing and connecting to machines by name instead of remembering their IP addresses.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
