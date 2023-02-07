package utils

import (
	"bytes"

	"os/exec"
	"strings"
)

func ExecuteShellCmd(command string, args ...string) ([]string, int, error) {
	// TODO validate input

	cmd := exec.Command(command, args...)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	if err != nil {
		Logger.Infoln(out.String())
		Logger.Errorln(err.Error())
	}

	var lines = DeleteEmpty(strings.Split(out.String(), "\n"))
	var statusCode = cmd.ProcessState.ExitCode()

	Logger.WithField("cmd", command+strings.Join(append([]string{" "}, args...), " ")).
		WithField("status_code", statusCode).
		WithField("output", lines).
		Traceln("executed shell command")

	return lines, statusCode, err
}
