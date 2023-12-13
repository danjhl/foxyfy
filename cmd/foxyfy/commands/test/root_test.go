package commands_test

import (
	"bytes"
	"foxyfy/cmd/foxyfy/commands"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Cmd = commands.Cmd

type mockCmd struct {
	name   string
	called bool
	args   []string
}

func (c *mockCmd) Name() string {
	return c.name
}

func (c *mockCmd) Help() string {
	return "help"
}

func (c *mockCmd) Execute(args []string) {
	c.called = true
	c.args = args
}

func TestRootCmdNoCommand(t *testing.T) {
	var out bytes.Buffer
	cmd := mockCmd{name: "some cmd"}
	args := []string{}

	rootCmd := commands.NewRootCmd(&out, []Cmd{&cmd})
	rootCmd.Execute(args)

	assert.Equal(t, "help\n\n", out.String())
}

func TestRootCmdCallHelp(t *testing.T) {
	var out bytes.Buffer
	cmd := mockCmd{name: "some cmd"}
	args := []string{"help"}

	rootCmd := commands.NewRootCmd(&out, []Cmd{&cmd})
	rootCmd.Execute(args)

	assert.Equal(t, "help\n\n", out.String())
}

func TestRootCmdCallCmd(t *testing.T) {
	var out bytes.Buffer
	updateCmd := mockCmd{name: "update"}
	args := []string{"update", "-b", "music"}

	rootCmd := commands.NewRootCmd(&out, []Cmd{&updateCmd})
	rootCmd.Execute(args)

	assert.Equal(t, true, updateCmd.called)
	assert.Equal(t, []string{"-b", "music"}, updateCmd.args)
}

func TestRootCmdUnknownCmd(t *testing.T) {
	var out bytes.Buffer
	rootCmd := commands.NewRootCmd(&out, []Cmd{})
	args := []string{"unknown"}

	rootCmd.Execute(args)

	assert.Equal(t, "Unknown command 'unknown'\n", out.String())
}
