package gitcmd

import (
	"bytes"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"log"
	"os/exec"
	"strings"
)

func ListCommitIDs() []string {
	cmd := exec.Command("git", "log", "--pretty=%h - %cn - %s")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Println(out.String())
		log.Fatal(err)
	}
	var lines = utils.DeleteEmpty(strings.Split(out.String(), "\n"))
	return lines
}
