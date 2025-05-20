package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xddprog/go_module_cli/internal/dependencies"
	"github.com/xddprog/go_module_cli/internal/modfile"
	"github.com/xddprog/go_module_cli/internal/repo"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go_modules_cli <git-repo-url>")
		return
	}

	target := os.Args[1]

	repoDir, err := repo.LoadRepo(target)
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(repoDir)

	moduleInfo, err := modfile.ParseGoMod(repoDir)
	if err != nil {
		log.Fatal(err)
	}

	updates, err := dependencies.CheckUpdates(repoDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Module: %s\n", moduleInfo.Name)
	fmt.Printf("Go Version: %s\n", moduleInfo.GoVersion)
	fmt.Println("Dependencies:")
	if len(updates) > 0 {
		fmt.Println("Updates available:")
		for _, update := range updates {
			fmt.Printf("Updating %s from %s to %s\n", update.Path, update.Current, update.Latest)
		}
	} else {
		fmt.Println("No updates available")
	}
}
