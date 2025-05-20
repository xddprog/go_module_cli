package dependencies

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)


type Dependency struct {
	Path    string
	Version string
	Update  struct {
		Version string
	}
}


type DependencyUpdate struct {
	Path    string
	Current string
	Latest  string
}


func CheckUpdates(repoPath string) ([]DependencyUpdate, error) {
	cmd := exec.Command("go", "list", "-u", "-m", "-json", "all")
	cmd.Dir = repoPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to check updates: %v\n%s", err, string(output))
	}

	var updates []DependencyUpdate
	decoder := json.NewDecoder(strings.NewReader(string(output)))
	
	for decoder.More() {
		var dep Dependency
		if err := decoder.Decode(&dep); err != nil {
			return nil, fmt.Errorf("failed to decode dependency: %s", err)
		}
		if dep.Update.Version != "" && dep.Update.Version != dep.Version {
			updates = append(updates, DependencyUpdate{
				Path:    dep.Path,
				Current: dep.Version,
				Latest:  dep.Update.Version,
			})
		}
	}

	return updates, nil
}