package command

import (
	"github.com/simonjohansson/cf-protocol/helpers"
)

type Plan struct {
	Cmds []Cmd
}

func (p Plan) Printable() []string {
	var commands []string

	for _, cmd := range p.Cmds {
		commands = append(commands, cmd.Printable())
	}

	return commands
}

func (p Plan) PrintPlan(logger helpers.Logging) {
	for _, command := range p.Printable() {
		logger.Info(command)
	}
}
