package commands_test

import (
	"bytes"
	"foxyfy/cmd/foxyfy/commands"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCmd struct {
	called bool
	args   []string
}

func (c *mockCmd) Execute(args []string) {
	c.called = true
	c.args = args
}

func TestRootCmdNoCmd(t *testing.T) {
	var out bytes.Buffer
	helpCmd := mockCmd{}
	args := []string{}

	rootCmd := commands.NewRootCmd(&out, &helpCmd, nil, nil)
	rootCmd.Execute(args)

	assert.Equal(t, true, helpCmd.called)
	assert.Equal(t, []string{}, helpCmd.args)
}

func TestRootCmdCallHelpCmd(t *testing.T) {
	var out bytes.Buffer
	helpCmd := mockCmd{}
	args := []string{"help"}

	rootCmd := commands.NewRootCmd(&out, &helpCmd, nil, nil)
	rootCmd.Execute(args)

	assert.Equal(t, true, helpCmd.called)
	assert.Equal(t, []string{}, helpCmd.args)
}

func TestRootCmdCallLsCmd(t *testing.T) {
	var out bytes.Buffer
	lsCmd := mockCmd{}
	args := []string{"ls", "-b", "music"}

	rootCmd := commands.NewRootCmd(&out, nil, &lsCmd, nil)
	rootCmd.Execute(args)

	assert.Equal(t, true, lsCmd.called)
	assert.Equal(t, []string{"-b", "music"}, lsCmd.args)
}

func TestRootCmdCallUpdateCmd(t *testing.T) {
	var out bytes.Buffer
	updateCmd := mockCmd{}
	args := []string{"update", "-b", "music"}

	rootCmd := commands.NewRootCmd(&out, nil, nil, &updateCmd)
	rootCmd.Execute(args)

	assert.Equal(t, true, updateCmd.called)
	assert.Equal(t, []string{"-b", "music"}, updateCmd.args)
}

func TestRootCmdUnknownCmd(t *testing.T) {
	var out bytes.Buffer
	rootCmd := commands.NewRootCmd(&out, nil, nil, nil)
	args := []string{"unknown"}

	rootCmd.Execute(args)

	assert.Equal(t, "Unknown command 'unknown'\n", out.String())
}
