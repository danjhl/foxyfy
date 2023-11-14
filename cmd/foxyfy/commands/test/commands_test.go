package commands_test

import (
	"bytes"
	"foxyfy/cmd/foxyfy/commands"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknownCmd(t *testing.T) {
	var out bytes.Buffer
	rootCmd := commands.NewRootCmd(&out)
	os.Args = []string{"bin", "unknown"}

	rootCmd.Execute()

	assert.Equal(t, "Unknown command 'unknown'\n", out.String())
}
