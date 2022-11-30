package gitcmd

import (
	"errors"
	"github.com/MauroMaia/gitmadeeasy/pkg/model"
	"github.com/MauroMaia/gitmadeeasy/pkg/utils"
	"regexp"
	"strings"
)

var findCurrentRegex, _ = regexp.Compile("^\\* (.*)")
var findRemotesRegex, _ = regexp.Compile("^remotes/(.*)/(.*)(.*)")

// TODO - fill the docs
func ListBranches() ([]model.Branch, error) {

	utils.Logger.WithField("func", "ListBranches").
		WithField("cmd", "git branch --no-color --list --all").
		Traceln("Listing branch's")

	output, exitCode, err := utils.ExecuteShellCmd("git", "branch", "--no-color", "--list", "--all")

	if err != nil || exitCode != 0 {
		return nil, errors.New(output[0])
	}

	var result []model.Branch

	for _, line := range output {

		if "" == line {
			continue
		}
		line = strings.Trim(line, " ")

		var isCurrent = findCurrentRegex.FindAllString(line, -1)

		if len(isCurrent) > 0 {
			var name = strings.Trim(isCurrent[0], " ")

			// FIXME - parse line to find if branch its local or remote or both
			var islocal = true

			var branch = model.NewBranch(name, islocal)

			result = append(result, branch)
			continue
		}

		var match = findRemotesRegex.FindAllString(line, -1)

		if len(match) > 0 {
			var name = strings.Trim(match[0], " ")

			// FIXME - parse line to find if branch its local or remote or both
			var islocal = false

			var branch = model.NewBranch(name, islocal)
			result = append(result, branch)
		}
	}

	return result, nil
}
