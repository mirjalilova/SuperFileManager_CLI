package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func FetchFile(url, newName, downloadDir string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the file:", err)
		return
	}
	defer resp.Body.Close()

	if newName == "" {
		newName = filepath.Base(url)
	}

	out, err := os.Create(filepath.Join(downloadDir, newName))
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error saving the file:", err)
		return
	}

	fmt.Println("File fetched successfully.")
}
