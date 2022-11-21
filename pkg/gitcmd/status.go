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

	return lines
}
