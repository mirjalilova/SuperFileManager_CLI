package filemanager

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"golang.org/x/exp/slog"
)

func UploadFile(srcPath, fileName string) error {
	file, err := os.Open(srcPath)
	if err != nil {
		slog.Error("Failed to open file", "error", err, "file", srcPath)
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	destDir := "./internal/files"
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		slog.Error("Failed to create directory", "error", err, "directory", destDir)
		return fmt.Errorf("failed to create directory: %w", err)
	}

	destPath := filepath.Join(destDir, fileName)

	out, err := os.Create(destPath)
	if err != nil {
		slog.Error("Failed to create destination file", "error", err, "file", destPath)
		return fmt.Errorf("could not create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		slog.Error("Failed to copy file", "error", err)
		return fmt.Errorf("could not copy file: %w", err)
	}

	slog.Info("File uploaded locally", "file", destPath)
	return nil
}