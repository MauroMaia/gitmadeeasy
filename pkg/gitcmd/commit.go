package gitcmd

import (
	"bytes"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"os"
	"os/exec"
	"strings"
)

func ListCommitIDs() []string {

	utils.Logger.Tracef("ListCommitIDs")

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

func Commit(message string, amend bool) ([]string, error) {

	utils.Logger.WithField("amend", amend).Tracef("commit")

	_ = os.Remove("/tmp/message")
	file, err := os.Create("/tmp/message")
	if err != nil {
		return nil, err
	}

	if _, err = file.WriteString(message); err != nil {
		return nil, err
	}

	var cmd *exec.Cmd

	if amend {
		cmd = exec.Command("git", "commit", "---amend", "-F", "/tmp/message")
	} else {
		cmd = exec.Command("git", "commit", "-F", "/tmp/message")
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	if err = cmd.Run(); err != nil {
		utils.Logger.Infoln(out.String())
		utils.Logger.Error(err)
		return nil, err
	}
	var lines = utils.DeleteEmpty(strings.Split(out.String(), "\n"))
	return lines, nil
}
