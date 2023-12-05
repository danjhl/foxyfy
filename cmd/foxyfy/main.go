package main

import (
	"fmt"
	"foxyfy/cmd/foxyfy/commands"
	"os"
	"os/exec"
)

type executor struct{}

func (e executor) RunFromPath(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	out, err := command.Output()
	if err != nil {
		fmt.Printf("exec: %v %v\n", cmd, args)
		fmt.Printf("err: %v\n %v\n", err, string(out[:]))
		return err
	}

	return nil
}

func main() {
	helpCmd := commands.NewHelpCmd(os.Stdout)
	lsCmd := commands.NewLsCmd(os.Stdout)
	updateCmd := commands.NewUpdateCmd(os.Stdout, executor{})
	rootCmd := commands.NewRootCmd(os.Stdout, helpCmd, lsCmd, updateCmd)
	rootCmd.Execute(os.Args[1:])
}
