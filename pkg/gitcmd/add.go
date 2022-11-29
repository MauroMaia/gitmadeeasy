package gitcmd

import (
	"bytes"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"os/exec"
	"strings"
)

func StageFile(filename string) []string {

	utils.Logger.WithField("filename", filename).
		WithField("func", "StageFile").
		WithField("cmd", "git add").
		Traceln("Adding git file")

	cmd := exec.Command("git", "add", filename)

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
