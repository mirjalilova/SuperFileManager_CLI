package cmd

import (
	"path/filepath"
	"superfilemanager/internal/minio"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var downloadCmd = &cobra.Command{
	Use:   "download [file]",
	Short: "Download a file from MinIO",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]

		minioClient, err := minio.MinIOConnect()
		if err != nil {
			slog.Error("Failed to connect to MinIO", "error", err)
			return
		}

		destPath := filepath.Join("./internal/files", fileName)

		err = minioClient.Download(fileName, destPath)
		if err != nil {
			slog.Error("Failed to download the file", "error", err, "file", fileName)
			return
		}

		slog.Info("File downloaded successfully", "file", fileName)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
