package utils

import (
	"os"
	"path/filepath"
)

func IsGitRepoDirectory() bool {
	leaf, _ := os.Getwd()

	// Traverse from leaf to root
	for {
		// Print the current directory
		// fmt.Println(leaf)

		// check if there is a .git directory in this leaf
		if stat, err := os.Stat(leaf + "/.git"); !os.IsNotExist(err) {
			if stat.IsDir() == true {
				return true
			}
		}

		// Get the parent directory
		parent := filepath.Dir(leaf)

		// Stop if the parent directory is the same as the current directory
		if parent == leaf {
			break
		}

		// Update the leaf directory
		leaf = parent
	}
	return false
}
