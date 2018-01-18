package command

import (
	"github.com/simonjohansson/cf-protocol/helpers"
	"code.cloudfoundry.org/cli/plugin"
	"os/exec"
)

type Executor interface {
	Execute(plan Plan) error
}

/////

type cfExecutor struct {
	cliConnection plugin.CliConnection
	logger        helpers.Logger
}

func NewCfExecutor(cliConnection plugin.CliConnection, logger helpers.Logger) cfExecutor {
	return cfExecutor{
		cliConnection: cliConnection,
		logger:        logger,
	}
}

func (e cfExecutor) Execute(plan Plan) error {
	for _, cmd := range plan.Cmds {
		e.logger.Info("About to execute: " + cmd.Printable())
		_, err := e.cliConnection.CliCommand(cmd.GetArgs()...)
		if err != nil {
			return err
		}
	}
	return nil
}

/////

type cliExecutor struct {
	cliConnection plugin.CliConnection
	logger        helpers.Logger
}

func NewCliExecutor(logger helpers.Logger) cliExecutor {
	return cliExecutor{
		logger: logger,
	}
}

func (e cliExecutor) Execute(plan Plan) error {
	for _, cmd := range plan.Cmds {
		e.logger.Info("About to execute: " + cmd.Printable())
		execCmd := exec.Command(cmd.GetArgs()[0], cmd.GetArgs()[1:]...)
		execCmd.Stdout = e.logger
		execCmd.Stderr = e.logger
		execCmd.Dir = cmd.GetDir()
		err := execCmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
