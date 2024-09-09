package cmd

import (
	"superfilemanager/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch a file from the internet",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		newName, _ := cmd.Flags().GetString("new-name")

		if url == "" {
			color.Red("Error: URL is required")
			slog.Error("URL is required for fetching the file")
			return
		}

		color.Cyan("Fetching file from URL: %s", url)
		err := utils.FetchFile(url, newName, "./internal/files")
		if err != nil {
			color.Red("Failed to fetch file: %v", err)
			slog.Error("Failed to fetch file", "url", url, "newName", newName, "error", err)
			return
		}

		color.Green("File fetched successfully")
		slog.Info("File fetched successfully", "url", url, "newName", newName)
	},
}

func init() {
	fetchCmd.Flags().StringP("url", "u", "", "URL of the file to fetch")
	fetchCmd.Flags().StringP("new-name", "n", "", "New name for the fetched file")
	rootCmd.AddCommand(fetchCmd)
}
