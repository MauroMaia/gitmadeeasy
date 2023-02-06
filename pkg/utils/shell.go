package utils

import (
	"bytes"

	"os/exec"
	"strings"
)

func ExecuteShellCmd(command string, args ...string) ([]string, int, error) {
	// TODO validate input

	Logger.WithField("cmd", command+strings.Join(append([]string{" "}, args...), " ")).
		WithField("args", args).
		Traceln("Executing shell command")

	cmd := exec.Command(command, args...)

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
