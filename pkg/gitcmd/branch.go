package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func ListBranches() string {
	// git  branch --no-color

	cmd := exec.Command("git", "branch --no-color --list --all")

	cmd.Stdin = strings.NewReader("and old falcon")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	return ""
}
