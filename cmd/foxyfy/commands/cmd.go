package commands

type Cmd interface {
	Name() string
	Help() string
	Execute(args []string)
}
