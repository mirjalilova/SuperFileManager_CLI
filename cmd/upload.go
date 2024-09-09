package cmd

import (
	"os"
	"path/filepath"
	"superfilemanager/utils/filemanager"
	"superfilemanager/internal/minio"

	"github.com/spf13/cobra"
	"github.com/fatih/color"
	"golang.org/x/exp/slog"
)

var uploadCmd = &cobra.Command{
	Use:   "upload [file]",
	Short: "Upload a file to MinIO",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		color.Cyan("Uploading file: %s", filePath)

		fileName := filepath.Base(filePath)

		err := filemanager.UploadFile(filePath, fileName)
		if err != nil {
			color.Red("Failed to upload file: %v", err)
			slog.Error("Failed to upload file", "error", err)
			return
		}

		color.Green("File uploaded successfully: %s", fileName)
		slog.Info("File uploaded successfully", "file", fileName)

		minioClient, err := minio.MinIOConnect()
		if err != nil {
			color.Red("Failed to connect to MinIO: %v", err)
			slog.Error("Failed to connect to MinIO", "error", err)
			return
		}

		filePath = filepath.Join("./internal/files", fileName)
		_, err = minioClient.Upload(fileName, filePath)
		if err != nil {
			color.Red("Failed to upload file to MinIO: %v", err)
			slog.Error("Failed to upload file to MinIO", "error", err)
			return
		}

		err = os.Remove(filePath)
		if err != nil {
			color.Red("Failed to delete the local file: %v", err)
			slog.Error("Failed to delete the local file", "error", err, "file", fileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
