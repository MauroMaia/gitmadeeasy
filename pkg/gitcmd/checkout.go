package gitcmd

import (
	"errors"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
)

// TODO - fill the docs
func CreateNewBranch(branchName string, push bool) error {

	utils.Logger.WithField("branchName", branchName).
		WithField("push", push).
		WithField("cmd", "git checkout -b").
		Traceln("Create local branch")

	// TODO validate input

	result, exitCode, err := utils.ExecuteShellCmd("git", "checkout", "-b", branchName)

	if err != nil || exitCode != 0 {
		return errors.New(result[0])
	}

	if push {
		utils.Logger.WithField("branchName", branchName).
			WithField("push", push).
			WithField("cmd", "git push --set-upstream origin").
			Traceln("push local branch to remote (hardcoded origin)")

		// FIXME - remote as origin should not be here hardcoded
		result, exitCode, err = utils.ExecuteShellCmd("git", "push", "--set-upstream", "origin", branchName)

		if err != nil || exitCode != 0 {
			return errors.New(result[0])
		}
	}

	return nil
}
