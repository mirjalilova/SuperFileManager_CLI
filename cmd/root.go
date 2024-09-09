package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var rootCmd = &cobra.Command{
	Use:   "superfile",
	Short: "SuperFile CLI application",
	Long:  `SuperFile is a CLI application for managing files`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Error executing command", "error", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(listCmd)
}
