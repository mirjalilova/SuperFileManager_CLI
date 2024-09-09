package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/exp/slog"
)

func FetchFile(url, newName, downloadDir string) error {
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Error fetching the file:", err)
		return err
	}
	defer resp.Body.Close()

	if newName == "" {
		newName = filepath.Base(url)
	}

	out, err := os.Create(filepath.Join(downloadDir, newName))
	if err != nil {
		slog.Error("Error creating the file:", err)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		slog.Error("Error saving the file:", err)
		return err
	}

	slog.Info("File fetched successfully.")
	return nil
}
