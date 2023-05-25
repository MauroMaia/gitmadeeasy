package gitcmd

import (
	"errors"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
)

// TODO - fill the docs
func Pull(prune bool, rebase bool, stash bool) ([]string, error) {

	var err error
	var exitCode int
	var result []string

	utils.Logger.
		WithField("prune", prune).
		WithField("rebase", rebase).
		WithField("stash", stash).
		Traceln("Get changes from remote repositories")

	args := []string{
		"pull",
	}

	if prune {
		args = append(args, "--prune")
	}

	if rebase {
		args = append(args, "--rebase")
	}

	result, exitCode, err = utils.ExecuteShellCmd("git", args...)
	if err != nil || exitCode != 0 {
		return nil, errors.New(result[1])
	}

	return result, nil
}
