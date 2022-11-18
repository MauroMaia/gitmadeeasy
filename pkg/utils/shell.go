package utils

import (
	"bytes"

	"os/exec"
	"strings"
)

func ExecuteShellCmd(command string, name ...string) ([]string, int, error) {
	// TODO validate input

	cmd := exec.Command(command, name...)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	if err != nil {
		Logger.Infoln(out.String())
		Logger.Errorln(err.Error())
	}

	//TODO - log outout
	var lines = DeleteEmpty(strings.Split(out.String(), "\n"))
	return lines, cmd.ProcessState.ExitCode(), err
}
