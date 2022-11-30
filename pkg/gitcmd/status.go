package gitcmd

import (
	"errors"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
)

// TODO - fill the docs
func ListFilesChanged() ([]string, error) {

	utils.Logger.WithField("func", "ListFilesChanged").
		WithField("cmd", "git status -s").
		Traceln("Get list of files changed")

	result, exitCode, err := utils.ExecuteShellCmd("git", "status", "-s")

	if err != nil || exitCode != 0 {
		return nil, errors.New(result[0])
	}

	/*
	 * TODO: create an object in memory with this data to allow better ui drawing
	 * Each line has the following format:
	 *		1º char staged - with major type of change (M: modified, R: renamed, D: Deleted, A/??: Added, etc?)
	 *		2º char local change - with major type of change (M: modified, R: renamed, D: Deleted, A/??: Added, etc?)
	 *		3º char+ path to the file changed
	 */

	return utils.DeleteEmpty(result), nil
}
