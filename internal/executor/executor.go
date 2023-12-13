package executor

import (
	"fmt"
	"os/exec"
)

type Executor interface {
	RunFromPath(cmd string, args ...string) error
}

type ExecExecutor struct{}

func (e ExecExecutor) RunFromPath(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	out, err := command.Output()
	if err != nil {
		fmt.Printf("exec: %v %v\n", cmd, args)
		fmt.Printf("err: %v\n %v\n", err, string(out[:]))
		return err
	}

	return nil
}
