package gitcmd

import (
	"bytes"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"os/exec"
	"strings"
)

func ListFilesChanged() []string {
	// git branch --no-color --list --all
	utils.Logger.Tracef("ListFilesChanged")
	cmd := exec.Command("git", "status", "-s")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	if err != nil {
		utils.Logger.Infoln(out.String())
		utils.Logger.Fatalln(err)
	}

	var lines = strings.Split(out.String(), "\n")

	/*
	 * TODO: create an object in memory with this data to allow better ui drawing
	 * Each line has the following format:
	 *		1º char staged - with major type of change (M: modified, R: renamed, D: Deleted, A/??: Added, etc?)
	 *		2º char local change - with major type of change (M: modified, R: renamed, D: Deleted, A/??: Added, etc?)
	 *		3º char+ path to the file changed
	 */

	return lines
}
