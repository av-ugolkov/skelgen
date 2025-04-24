package actions

import (
	"os/exec"
)

func ExecCmdInDir(dir, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Dir = dir

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}

func ExecCmd(cmd string, args ...string) error {
	return ExecCmdInDir(cmd, ".", args...)
}
