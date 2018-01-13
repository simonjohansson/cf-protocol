package promote

import (
	"code.cloudfoundry.org/cli/plugin"
	. "github.com/simonjohansson/cf-protocol/command"
	. "github.com/simonjohansson/cf-protocol/helpers"
	"flag"
)

type Promote struct {
	cliConnection plugin.CliConnection
	logger        Logger
}

func NewPromote(cliConnection plugin.CliConnection, logger Logger) Promote {
	return Promote{
		cliConnection,
		logger,
	}
}

func (p Promote) Promote(manifestPath string) Plan {
	return Plan{
		Cmds: []Cmd{
			CfApps{},
		},
	}
}

func (p Promote) RunPromote(args []string) error {
	flagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	manifestPath := flagSet.String("manifest", "", "Path to the manifest")
	err := ParseArgs(p.logger, flagSet, args)
	if err != nil {
		return err
	}

	p.logger.Info("Starting promote")
	plan := p.Promote(*manifestPath)

	err = plan.ExecutePlan(p.cliConnection, p.logger)
	if err != nil {
		return err
	}

	return nil
}
