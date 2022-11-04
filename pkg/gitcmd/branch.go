package gitcmd

import (
	"bytes"
	"github.com/MauroMaia/gitmadeeasy/pkg/model"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

var findCurrentRegex, _ = regexp.Compile("^\\* (.*)")
var findRemotesRegex, _ = regexp.Compile("^remotes/(.*)/(.*)(.*)")

func ListBranches() []model.Branch {
	// git branch --no-color --list --all

	cmd := exec.Command("git", "branch", "--no-color", "--list", "--all")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Println(out.String())
		log.Fatal(err)
	}

	var result []model.Branch

	var lines = strings.Split(out.String(), "\n")
	for _, line := range lines {

		if "" == line {
			continue
		}
		line = strings.Trim(line, " ")

		var isCurrent = findCurrentRegex.FindAllString(line, -1)

		if len(isCurrent) > 0 {
			var name = strings.Trim(isCurrent[0], " ")
			var islocal = true
			var branch = model.NewBranch(name, islocal)
			result = append(result, branch)
			continue
		}

		var match = findRemotesRegex.FindAllString(line, -1)

		if len(match) > 0 {
			var name = strings.Trim(match[0], " ")
			var islocal = false
			var branch = model.NewBranch(name, islocal)
			result = append(result, branch)
		}
	}

	return result
}
