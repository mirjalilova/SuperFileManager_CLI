package utils

import (
    "fmt"
    "os/exec"
)

func CloneRepository(url, newName, cloneDir string) {
    cmd := exec.Command("git", "clone", url)
    if newName != "" {
        cmd.Args = append(cmd.Args, newName)
    }
    cmd.Dir = cloneDir
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Repository cloned successfully.")
}
