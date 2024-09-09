package utils

import (
	"os/exec"

	"golang.org/x/exp/slog"
)

func CloneRepository(url, newName, cloneDir string) error {
    cmd := exec.Command("git", "clone", url)
    if newName != "" {
        cmd.Args = append(cmd.Args, newName)
    }
    cmd.Dir = cloneDir
    err := cmd.Run()
    if err != nil {
        slog.Error("Error:", err)
        return err
    }

    slog.Info("Repository cloned successfully.")
    return nil
}
