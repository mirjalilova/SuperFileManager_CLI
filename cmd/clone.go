package cmd

import (
	"superfilemanager/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		newName, _ := cmd.Flags().GetString("new-name")

		if url == "" {
			color.Red("Error: URL is required")
			slog.Error("URL is required for cloning the repository")
			return
		}

		color.Cyan("Cloning repository from URL: %s", url)
		err := utils.CloneRepository(url, newName, "./internal/repos")
		if err != nil {
			color.Red("Failed to clone repository: %v", err)
			slog.Error("Failed to clone repository", "url", url, "newName", newName, "error", err)
			return
		}

		color.Green("Repository cloned successfully")
		slog.Info("Repository cloned successfully", "url", url, "newName", newName)
	},
}

func init() {
	cloneCmd.Flags().StringP("url", "u", "", "URL of the repository to clone")
	cloneCmd.Flags().StringP("new-name", "n", "", "New name for the cloned repository")
	rootCmd.AddCommand(cloneCmd)
}
