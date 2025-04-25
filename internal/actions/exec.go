package actions

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecCmdInDir(dir, cmd string) error {
	commands := strings.Split(cmd, " ")
	c := exec.Command(commands[0], commands[1:]...)
	c.Dir = dir

	err := c.Run()
	if err != nil {
		return fmt.Errorf("couldn't exec command [%s] into folder [%s]: %v", cmd, dir, err)
	}

	return nil
}

func ExecCmd(cmd string) error {
	return ExecCmdInDir(cmd, ".")
}
