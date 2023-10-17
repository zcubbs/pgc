package main

import (
	"fmt"
	"os/exec"
)

func restartPostgres(command string) error {
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error restarting postgres: %w", err)
	}
	return nil
}
