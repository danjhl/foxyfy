package main

import (
	"foxyfy/cmd/foxyfy/commands"
	"os"
)

func main() {
	rootCmd := commands.NewRootCmd(os.Stdout)
	rootCmd.Execute()
}
