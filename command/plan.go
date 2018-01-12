package command

import (
	"strings"
	"github.com/simonjohansson/cf-protocol/helpers"
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
)

type Cmd interface {
	ExecuteCmd(cliConnection plugin.CliConnection, logger helpers.Logger) ([]string, error)
	Printable() string
}

type CfCmd struct {
	Args []string
}

func (c CfCmd) ExecuteCmd(cliConnection plugin.CliConnection, logger helpers.Logger) ([]string, error) {
	logger.Info("About to execute: " + c.Printable())
	return cliConnection.CliCommand(c.Args...)
}

func (c CfCmd) Printable() string {
	return "cf " + strings.Join(c.Args, " ")
}

type CfApps struct{}

func (c CfApps) ExecuteCmd(cliConnection plugin.CliConnection, logger helpers.Logger) ([]string, error) {
	apps, err := cliConnection.GetApps()
	if err != nil {
		return []string{}, err
	}

	for _, app := range apps {
		fmt.Println(app.Name)
	}
	return []string{}, nil
}

func (c CfApps) Printable() string {
	return "cf apps"
}

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

func (p Plan) PrintPlan(logger helpers.Logger) {
	for _, command := range p.Printable() {
		logger.Info("\t" + command)
	}
}

func (p Plan) ExecutePlan(cliConnection plugin.CliConnection, logger helpers.Logger) error {
	for _, cmd := range p.Cmds {
		_, err := cmd.ExecuteCmd(cliConnection, logger)
		if err != nil {
			return err
		}
	}

	return nil
}
