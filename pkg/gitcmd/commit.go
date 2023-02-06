package gitcmd

import (
	"errors"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"os"
)

// TODO - fill the docs
func ListCommits() ([]string, error) {

	utils.Logger.WithField("cmd", "git log --pretty=%h - %cn - %s").
		Traceln("Listing commits")

	result, exitCode, err := utils.ExecuteShellCmd("git", "log", "--pretty=%h - %cn - %s")

	if err != nil || exitCode != 0 {
		return nil, errors.New(result[0])
	}

	return utils.DeleteEmpty(result), nil
}

// TODO - fill the docs
func Commit(message string, amend bool) ([]string, error) {

	utils.Logger.WithField("amend", amend).
		WithField("message", message).
		WithField("cmd", "git commit [--amend] -F /tmp/message").
		Traceln("commit files")

	_ = os.Remove("/tmp/message")
	file, err := os.Create("/tmp/message")
	if err != nil {
		return nil, err
	}

	if _, err = file.WriteString(message); err != nil {
		return nil, err
	}

	var result []string
	var exitCode int

	if amend {
		result, exitCode, err = utils.ExecuteShellCmd("git", "commit", "---amend", "-F", "/tmp/message")
	} else {
		result, exitCode, err = utils.ExecuteShellCmd("git", "commit", "-F", "/tmp/message")
	}

	if err != nil || exitCode != 0 {
		return nil, errors.New(result[0])
	}

	return utils.DeleteEmpty(result), nil
}
