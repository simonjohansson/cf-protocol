package push

import (
	. "github.com/simonjohansson/cf-protocol/helpers"
	. "github.com/simonjohansson/cf-protocol/command"
	"fmt"
	"code.cloudfoundry.org/cli/plugin"
)

type Push struct {
	manifestReader ManifestReader
	cliConnection  plugin.CliConnection
	options        Options
}

func NewPush(manifestReader ManifestReader, cliConnection plugin.CliConnection, options Options) Push {
	return Push{
		manifestReader,
		cliConnection,
		options,
	}
}

func (p Push) appName(baseName string, postfix string) string {
	return fmt.Sprintf("%s-%s", baseName, postfix)
}

func (p Push) testRoute(baseName string) (string, error) {
	currentSpace, err := p.cliConnection.GetCurrentSpace()
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%s-%s-%s", baseName, currentSpace.Name, "test"), nil
}

func (p Push) PushPlan() (Plan, error) {
	application, err := p.manifestReader.Read(p.options.ManifestPath)
	if err != nil {
		return Plan{}, err
	}

	appName := p.appName(application.Name, p.options.Postfix)
	appHostname, err := p.testRoute(application.Name)
	if err != nil {
		return Plan{}, err
	}
	cmd := CfCmd{
		[]string{"push", appName, "-f", p.options.ManifestPath, "-n", appHostname, "-d", p.options.Domain},
	}

	return Plan{[]Cmd{cmd}}, nil
}
