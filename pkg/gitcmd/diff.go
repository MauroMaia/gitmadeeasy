package gitcmd

import (
	"errors"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
)

// TODO - fill the docs
func GetDiffPatch() ([]string, error) {

	utils.Logger.WithField("func", "GetDiffPatch").
		WithField("cmd", "git diff -p").
		Traceln("Get diff from remote version in patch format")

	result, exitCode, err := utils.ExecuteShellCmd("git", "diff", "-p")

	if err != nil || exitCode != 0 {
		return nil, errors.New(result[0])
	}

	return result, nil
}

// TODO - fill the docs
func GetDiffPatchForFile(filename string) ([]string, error) {

	utils.Logger.WithField("func", "GetDiffPatchForFile").
		WithField("filename", filename).
		WithField("cmd", "git diff -p --").
		Traceln("Get diff from remote version in patch format")

	result, exitCode, err := utils.ExecuteShellCmd("git", "diff", "-p", "--", filename)

	if err != nil || exitCode != 0 {
		return nil, errors.New(result[0])
	}

	return result, nil
}
