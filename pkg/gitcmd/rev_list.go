package gitcmd

import (
	"errors"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
)

// GetNrOfCommitsAhead func TODO - fill the docs
func GetNrOfCommitsAhead() (string, error) {

	currentBranch, err := GetCurrentBranchName()
	if err != nil {
		return "", err
	}

	// FIXME - This will not work with multiple origins
	currentBranch = "origin/" + currentBranch

	utils.Logger.WithField("currentBranch", currentBranch).
		WithField("cmd", `git rev-list ${currentBranch}..HEAD --count`).
		Traceln("Adding git file")

	result, exitCode, err := utils.ExecuteShellCmd("git", "rev-list", currentBranch+"..HEAD", "--count")

	if err != nil || exitCode != 0 {
		return "", errors.New(result[0])
	}

	return result[0], nil
}

// GetNrOfCommitsBehind func TODO - fill the docs
func GetNrOfCommitsBehind() (string, error) {

	currentBranch, err := GetCurrentBranchName()
	if err != nil {
		return "", err
	}

	// FIXME - This will not work with multiple origins
	currentBranch = "origin/" + currentBranch

	utils.Logger.WithField("currentBranch", currentBranch).
		WithField("cmd", `git rev-list HEAD..${currentBranch} --count`).
		Traceln("Adding git file")

	result, exitCode, err := utils.ExecuteShellCmd("git", "rev-list", "HEAD.."+currentBranch, "--count")

	if err != nil || exitCode != 0 {
		return "", errors.New(result[0])
	}

	return result[0], nil
}
