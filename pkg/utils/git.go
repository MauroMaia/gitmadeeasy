package utils

import "os"

func IsGitRepoDirectory() bool {
	// FIXME - must check recursively
	if stat, err := os.Stat(".git"); !os.IsNotExist(err) {
		return stat.IsDir()
	}
	return false
}
