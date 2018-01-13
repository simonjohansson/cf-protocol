package delete

import (
	. "github.com/simonjohansson/cf-protocol/helpers"
	. "github.com/simonjohansson/cf-protocol/command"
	"fmt"
	"code.cloudfoundry.org/cli/plugin"
	"flag"
)

type Delete struct {
	cliConnection  plugin.CliConnection
	manifestReader ManifestReader
	logger         Logger
}

func NewDelete(cliConnection plugin.CliConnection, manifestReader ManifestReader, logger Logger) Delete {
	return Delete{
		cliConnection,
		manifestReader,
		logger,
	}
}

func (d Delete) appName(baseName string, postfix string) string {
	return fmt.Sprintf("%s-%s", baseName, postfix)
}

func (d Delete) DeletePlan(manifestPath string, postfix string) (Plan, error) {
	application, err := d.manifestReader.Read(manifestPath)
	if err != nil {
		return Plan{}, err
	}

	appName := d.appName(application.Name, postfix)
	return Plan{[]Cmd{
		CfCmd{[]string{"delete", appName, "-f", "-r"}},
	}}, nil
}

func (d Delete) RunDelete(args []string) error {
	flagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	manifestPath := flagSet.String("manifest", "", "Path to the manifest")
	postfix := flagSet.String("postfix", "", "Postfix to use push")
	err := ParseArgs(d.logger, flagSet, args)
	if err != nil {
		return err
	}

	plan, err := d.DeletePlan(*manifestPath, *postfix)
	if err != nil {
		return err
	}

	d.logger.Info("Execution plan")
	plan.PrintPlan(d.logger)

	d.logger.Info("Executing")
	err = plan.ExecutePlan(d.cliConnection, d.logger)
	if err != nil {
		return err
	}

	d.logger.Info("Delete succeeded!")
	return nil
}
