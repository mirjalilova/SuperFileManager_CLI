package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"superfilemanager/utils"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		newName, _ := cmd.Flags().GetString("new-name")

		if url == "" {
			fmt.Println("Error: URL is required")
			return
		}

		utils.CloneRepository(url, newName, "./internal/repos")
	},
}

func init() {
	cloneCmd.Flags().StringP("url", "u", "", "URL of the repository to clone")
	cloneCmd.Flags().StringP("new-name", "n", "", "New name for the cloned repository")
	rootCmd.AddCommand(cloneCmd)
}
