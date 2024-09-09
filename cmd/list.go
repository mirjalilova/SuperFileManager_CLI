package cmd

import (
	"superfilemanager/internal/minio"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
	"golang.org/x/exp/slog"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all files in MinIO",
	Run: func(cmd *cobra.Command, args []string) {
		minioClient, err := minio.MinIOConnect()
		if err != nil {
			color.Red("Failed to connect to MinIO: %v", err)
			slog.Error("Failed to connect to MinIO", "error", err)
			return
		}

		files, err := minioClient.ListFiles()
		if err != nil {
			color.Red("Failed to list files: %v", err)
			slog.Error("Failed to list files", "error", err)
			return
		}

		color.Cyan("Files in MinIO (count: %d)", len(files))
		for i, file := range files {
			color.Yellow("File [%d]: %s", i+1, file)
		}

		slog.Info("Listed files from MinIO", "count", len(files))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
