package modfile

import (
	"os"
	"path/filepath"
	"strings"
)


type ModuleInfo struct {
	Name string
	GoVersion string
}


func ParseGoMod(repoPath string) (*ModuleInfo, error) {
	filePath := filepath.Join(repoPath, "go.mod")

    content, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    info := &ModuleInfo{}
    lines := strings.Split(string(content), "\n")

    for _, line := range lines {
        line = strings.TrimSpace(line)
        if strings.HasPrefix(line, "module ") {
            info.Name = strings.TrimPrefix(line, "module ")
        } else if strings.HasPrefix(line, "go ") {
            info.GoVersion = strings.TrimPrefix(line, "go ")
        }
    }

    if info.Name == "" {
        info.Name = "unknown"
    }

    if info.GoVersion == "" {
        info.GoVersion = "unknown"
    }

    return info, nil
}
