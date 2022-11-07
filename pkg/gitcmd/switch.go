package gitcmd

import (
	"bytes"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"log"
	"os/exec"
	"strings"
)

func CreateNewBranch(name string) {
	// TODO validate input
	
	// git switch <name> ??
	cmd := exec.Command("git", "switch", name)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Println(out.String())
		log.Fatal(err)
	}
	//TODO - log outout
	// var lines = utils.DeleteEmpty(strings.Split(out.String(), "\n"))
	_ = utils.DeleteEmpty(strings.Split(out.String(), "\n"))
}
