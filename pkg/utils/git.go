package utils

import "os"

func IsGitRepoDirectory() bool {
	if stat, err := os.Stat(".git"); !os.IsNotExist(err) {
		return stat.IsDir()
	}
	return false
}
