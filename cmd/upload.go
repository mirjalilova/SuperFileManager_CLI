package cmd

import (
	"os"
	"path/filepath"
	"superfilemanager/internal/filemanager"
	"superfilemanager/internal/minio"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var uploadCmd = &cobra.Command{
	Use:   "upload [file]",
	Short: "Upload a file to MinIO",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		slog.Info("File is being uploaded", "file", filePath)

		fileName := filepath.Base(filePath)

		err := filemanager.UploadFile(filePath, fileName)
		if err != nil {
			slog.Error("Failed to upload file", "error", err)
			return
		}

		slog.Info("File uploaded successfully", "file", fileName)

		minioClient, err := minio.MinIOConnect()
		if err != nil {
			slog.Error("Failed to connect to MinIO", "error", err)
			return
		}

		filePath = filepath.Join("./internal/files", fileName)
		_, err = minioClient.Upload(fileName, filePath)
		if err != nil {
			slog.Error("Failed to upload file to MinIO", "error", err)
			return
		}

		err = os.Remove(filePath)
		if err != nil {
			slog.Error("Failed to delete the local file", "error", err, "file", fileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
