package promote

import (
	"code.cloudfoundry.org/cli/plugin"
	. "github.com/simonjohansson/cf-protocol/command"
	. "github.com/simonjohansson/cf-protocol/helpers"
	"flag"
)

func Promote(manifestPath string) Plan {
	return Plan{
		Cmds: []Cmd{
			CfApps{},
		},
	}
}

func RunPromote(cliConnection plugin.CliConnection, logger Logger, args []string) error {
	flagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	manifestPath := flagSet.String("manifest", "", "Path to the manifest")
	err := ParseArgs(logger, flagSet, args)
	if err != nil {
		return err
	}

	logger.Info("Starting promote")
	plan := Promote(*manifestPath)

	err = plan.ExecutePlan(cliConnection, logger)
	if err != nil {
		return err
	}

	return nil
}
