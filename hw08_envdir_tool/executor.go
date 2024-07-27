package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

var ErrEmptyCmd = errors.New("command is empty")

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	// Place your code here.
	if len(cmd) == 0 {
		log.Println(ErrEmptyCmd)
		return 1
	}

	for eVar, eVal := range env {
		if eVal.NeedRemove {
			err := os.Unsetenv(eVar)
			if err != nil {
				return 1
			}
			continue
		}

		err := os.Setenv(eVar, eVal.Value)
		if err != nil {
			return 1
		}
	}

	commandName := cmd[0]
	args := cmd[1:]
	proc := exec.Command(commandName, args...)
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	proc.Stdin = os.Stdin

	if err := proc.Run(); err != nil {
		return 1
	}

	return
}
