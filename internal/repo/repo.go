package repo

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)


func LoadRepo(target string) (string, error) {
	if !(strings.HasPrefix(target, "https://") || strings.HasPrefix(target, "git@")) {
		return "", fmt.Errorf("invalid target: %s", target)
	}

	tempDir, err := os.MkdirTemp("", "go-mod-updater")
	if err != nil {
		return "", err
	}

	cmd := exec.Command("git", "clone", target, tempDir)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return tempDir, nil
}