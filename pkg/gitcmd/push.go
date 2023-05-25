package gitcmd

import (
	"errors"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
)

// TODO - fill the docs
func Push(force bool) ([]string, error) {

	var err error
	var exitCode int
	var result []string

	utils.Logger.
		WithField("force", force).
		Traceln("Send changes to remote repositories")

	if force {
		result, exitCode, err = utils.ExecuteShellCmd("git", "push", "-f")
	} else {
		result, exitCode, err = utils.ExecuteShellCmd("git", "push")
	}

	if err != nil || exitCode != 0 {
		return nil, errors.New(result[1])
	}

	return result, nil
}
