package cmd

import (
    "superfilemanager/utils"
    "fmt"
    "github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
    Use:   "fetch",
    Short: "Fetch a file from the internet",
    Run: func(cmd *cobra.Command, args []string) {
        url, _ := cmd.Flags().GetString("url")
        newName, _ := cmd.Flags().GetString("new-name")

        if url == "" {
            fmt.Println("Error: URL is required")
            return
        }

        utils.FetchFile(url, newName, "./internal/files")
    },
}

func init() {
    fetchCmd.Flags().StringP("url", "u", "", "URL of the file to fetch")
    fetchCmd.Flags().StringP("new-name", "n", "", "New name for the fetched file")
    rootCmd.AddCommand(fetchCmd)
}
