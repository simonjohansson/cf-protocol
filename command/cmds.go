package command

import (
	"strings"
)

type Cmd interface {
	Printable() string
	GetArgs() []string
}

/////

type CfCmd struct {
	Args []string
}

func (c CfCmd) Printable() string {
	return "cf " + strings.Join(c.Args, " ")
}

func (c CfCmd) GetArgs() []string {
	return c.Args
}

/////

type CliCmd struct {
	Args []string
}

func (c CliCmd) Printable() string {
	return strings.Join(c.Args, " ")
}

func (c CliCmd) GetArgs() []string {
	return c.Args
}
