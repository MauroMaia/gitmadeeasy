package gitcmd

import (
	"errors"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
)

// TODO - fill the docs
func Fetch() ([]string, error) {

	utils.Logger.Traceln("Get changes from remote repositories")

	result, exitCode, err := utils.ExecuteShellCmd("git", "fetch", "--all")

	if err != nil || exitCode != 0 {
		return nil, errors.New(result[0])
	}

	return result, nil
}
