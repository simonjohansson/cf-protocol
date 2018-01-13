package push

import (
	. "github.com/simonjohansson/cf-protocol/helpers"
	. "github.com/simonjohansson/cf-protocol/command"
	"fmt"
)

type Push struct {
	manifestReader ManifestReader
	options        Options
}

func NewPush(manifestReader ManifestReader, options Options) Push {
	return Push{
		manifestReader,
		options,
	}
}

func (p Push) appName(baseName string, postfix string) string {
	return fmt.Sprintf("%s-%s", baseName, postfix)
}

func (p Push) PushPlan() (Plan, error) {
	application, err := p.manifestReader.Read(p.options.ManifestPath)
	if err != nil {
		return Plan{}, err
	}

	appName := p.appName(application.Name, p.options.Postfix)
	cmd := CfCmd{
		[]string{"push", appName, "-f", p.options.ManifestPath, "-n", appName, "-d", p.options.Domain},
	}

	return Plan{[]Cmd{cmd}}, nil
}
