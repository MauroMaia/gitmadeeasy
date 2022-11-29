package gitcmd

import (
	"bytes"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"os/exec"
	"strings"
)

func GetDiffPatch() []string {

	utils.Logger.WithField("func", "GetDiffPatch").
		WithField("cmd", "git diff -p").
		Traceln("Get diff from remote version in patch format")

	cmd := exec.Command("git", "diff", "-p")

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
func GetDiffPatchForFile(filename string) []string {

	utils.Logger.WithField("func", "GetDiffPatchForFile").
		WithField("filename", filename).
		WithField("cmd", "git diff -p --").
		Traceln("Get diff from remote version in patch format")

	cmd := exec.Command("git", "diff", "-p", "--", filename)

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
