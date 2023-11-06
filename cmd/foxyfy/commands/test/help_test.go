package commands_test

import (
	"bytes"
	"foxyfy/cmd/foxyfy/commands"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelpCmd(t *testing.T) {
	var out bytes.Buffer
	rootCmd := commands.NewRootCmd(&out)
	os.Args = []string{"bin", "help"}
	rootCmd.Execute()

	assert.Equal(t, expected(), out.String())
}

func TestHelpCmdIsDefault(t *testing.T) {
	var out bytes.Buffer
	rootCmd := commands.NewRootCmd(&out)
	os.Args = []string{"bin"}
	rootCmd.Execute()

	assert.Equal(t, expected(), out.String())
}

func expected() string {
	return "\n" +
		"Usage: [CMD] [FLAGS]\n\n" +
		"ls -b [bookmark]                     List youtube bookmarks in parent [bookmark]\n\n"
}
