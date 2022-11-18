package gitcmd

import (
	"bytes"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"os/exec"
	"strings"
)

func ListCommitIDs() []string {
	cmd := exec.Command("git", "log", "--pretty=%h - %cn - %s")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		utils.Logger.Infoln(out.String())
		utils.Logger.Fatalln(err)
	}
	var lines = utils.DeleteEmpty(strings.Split(out.String(), "\n"))
	return lines
}
