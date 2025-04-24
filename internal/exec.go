package internal

import (
	"os/exec"
)

func runInDir(dir, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Dir = dir

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}

func run(cmd string, args ...string) error {
	return runInDir(cmd, ".", args...)
}
