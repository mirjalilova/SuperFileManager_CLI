package cmd

import (
	"superfilemanager/internal/minio"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all files in MinIO",
	Run: func(cmd *cobra.Command, args []string) {
		minioClient, err := minio.MinIOConnect()
		if err != nil {
			slog.Error("Failed to connect to MinIO", "error", err)
			return
		}

		files, err := minioClient.ListFiles()
		if err != nil {
			slog.Error("Failed to list files", "error", err)
			return
		}

		slog.Info("Files in MinIO", "count", len(files))
		for i, file := range files {
			slog.Info("File", "index", i+1, "file", file)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
