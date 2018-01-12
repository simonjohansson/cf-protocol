package push

import (
	. "github.com/simonjohansson/cf-protocol/helpers"
	. "github.com/simonjohansson/cf-protocol/command"
	"fmt"
	"code.cloudfoundry.org/cli/plugin"
	"flag"
)

func appName(baseName string, postfix string) string {
	return fmt.Sprintf("%s-%s", baseName, postfix)
}

func PushPlan(manifestPath string, postfix string, domain string, logger Logger, manifestReader ManifestReader) (Plan, error) {
	application, err := manifestReader.Read(manifestPath)
	if err != nil {
		return Plan{}, err
	}

	appName := appName(application.Name, postfix)
	cmd := CfCmd{
		[]string{"push", appName, "-f", manifestPath, "-n", appName, "-d", domain},
	}

	return Plan{[]Cmd{cmd}}, nil
}

func RunPush(cliConnection plugin.CliConnection, logger Logger, args []string) error {
	flagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	manifestPath := flagSet.String("manifest", "", "Path to the manifest")
	postfix := flagSet.String("postfix", "", "Postfix to use push")
	domain := flagSet.String("domain", "", "Domain to use when pushing")
	err := ParseArgs(logger, flagSet, args)
	if err != nil {
		return err
	}

	plan, err := PushPlan(*manifestPath, *postfix, *domain, logger, NewManifestReader())
	if err != nil {
		return err
	}

	logger.Info("Execution plan")
	plan.PrintPlan(logger)

	logger.Info("Executing")
	err = plan.ExecutePlan(cliConnection, logger)
	if err != nil {
		return err
	}

	logger.Info("Push succeeded!")
	return nil
}
