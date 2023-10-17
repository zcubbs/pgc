package pgc

import (
	"fmt"
	"os/exec"
)

func RestartPostgres(command string) error {
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error restarting postgres: %w", err)
	}
	return nil
}
