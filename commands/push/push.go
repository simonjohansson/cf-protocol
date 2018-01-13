package push

import (
	. "github.com/simonjohansson/cf-protocol/helpers"
	. "github.com/simonjohansson/cf-protocol/command"
	"fmt"
	"code.cloudfoundry.org/cli/plugin"
	"flag"
)

type Push struct {
	cliConnection  plugin.CliConnection
	manifestReader ManifestReader
	logger         Logger
}

func NewPush(cliConnection plugin.CliConnection, manifestReader ManifestReader, logger Logger) Push {
	return Push{
		cliConnection,
		manifestReader,
		logger,
	}
}

func (p Push) appName(baseName string, postfix string) string {
	return fmt.Sprintf("%s-%s", baseName, postfix)
}

func (p Push) PushPlan(manifestPath string, postfix string, domain string) (Plan, error) {
	application, err := p.manifestReader.Read(manifestPath)
	if err != nil {
		return Plan{}, err
	}

	appName := p.appName(application.Name, postfix)
	cmd := CfCmd{
		[]string{"push", appName, "-f", manifestPath, "-n", appName, "-d", domain},
	}

	return Plan{[]Cmd{cmd}}, nil
}

func (p Push) RunPush(args []string) error {
	flagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	manifestPath := flagSet.String("manifest", "", "Path to the manifest")
	postfix := flagSet.String("postfix", "", "Postfix to use push")
	domain := flagSet.String("domain", "", "Domain to use when pushing")
	err := ParseArgs(p.logger, flagSet, args)
	if err != nil {
		return err
	}

	plan, err := p.PushPlan(*manifestPath, *postfix, *domain)
	if err != nil {
		return err
	}

	p.logger.Info("Execution plan")
	plan.PrintPlan(p.logger)

	p.logger.Info("Executing")
	err = plan.ExecutePlan(p.cliConnection, p.logger)
	if err != nil {
		return err
	}

	p.logger.Info("Push succeeded!")
	return nil
}
