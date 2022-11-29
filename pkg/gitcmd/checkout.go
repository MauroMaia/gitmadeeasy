package gitcmd

import (
	"errors"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
)

func CreateNewBranch(name string, push bool) (bool, error) {

	utils.Logger.WithField("func", "ListFilesChanged").
		WithField("cmd", "git status -s").
		Traceln("Get list of files changed")

	// TODO validate input

	// git switch <name> ?? vs checkout -b <name>
	// TODO git push --set-upstream origin wip-create-menu
	result, exitCode, err := utils.ExecuteShellCmd("git", "checkout", "-b", name)

	returnVal := err != nil && exitCode != 0
	if returnVal {
		return returnVal, errors.New(result[0])
	}
	// TODO validate shell output
	return returnVal, nil
}
