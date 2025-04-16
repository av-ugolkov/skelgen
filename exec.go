package main

import (
	"os/exec"
)

func runCmdInRootDir(cmd string, args ...string) error {
	return runCmd(".", cmd, args...)
}

func runCmd(dir, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Dir = dir

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}
