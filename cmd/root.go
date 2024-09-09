package cmd

import (
	"os"
	"log"
	"path/filepath"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
	"github.com/fatih/color"
)

var rootCmd = &cobra.Command{
	Use:   "superfile",
	Short: "SuperFile CLI application",
	Long:  `SuperFile is a CLI application for managing files`,
}

func Execute() {
	logFile, err := os.OpenFile(filepath.Join("internal", "logs", "info.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		color.Red("Error creating log file: %v", err)
		os.Exit(1)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	if err := rootCmd.Execute(); err != nil {
		color.Red("Error executing command: %v", err)
		slog.Error("Error executing command", "error", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(cloneCmd)
	rootCmd.AddCommand(fetchCmd)
}
