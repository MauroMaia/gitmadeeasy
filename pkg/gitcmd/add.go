package gitcmd

import (
	"errors"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"os"
)

// StageFile func encapsulate the logic to prepare files to be submitted
// It requires a filename with the path to the file
func StageFile(filename string) ([]string, error) {

	utils.Logger.WithField("filename", filename).
		WithField("func", "StageFile").
		WithField("cmd", "git add").
		Traceln("Adding git file")

	// check if filename still exists
	if _, err := os.Stat(filename); err != nil {
		return nil, err
	}

	result, exitCode, err := utils.ExecuteShellCmd("git", "add", filename)

	if err != nil || exitCode != 0 {
		return nil, errors.New(result[0])
	}

	return result, nil
}
