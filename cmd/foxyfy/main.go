package main

import (
	"foxyfy/cmd/foxyfy/commands"
	"foxyfy/internal/executor"
	"os"
)

func main() {
	updateCmd := commands.NewUpdateCmd(os.Stdout, executor.ExecExecutor{})
	rootCmd := commands.NewRootCmd(os.Stdout, []commands.Cmd{updateCmd})
	rootCmd.Execute(os.Args[1:])
}
