package exec

import (
	"os/exec"
)

func RunInDir(dir, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Dir = dir

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}

func Run(cmd string, args ...string) error {
	return RunInDir(cmd, ".", args...)
}
