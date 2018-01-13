package delete

import (
	. "github.com/simonjohansson/cf-protocol/helpers"
	. "github.com/simonjohansson/cf-protocol/command"
	"fmt"
)

type Delete struct {
	manifestReader ManifestReader
	options        Options
}

func NewDelete(manifestReader ManifestReader, options Options) Delete {
	return Delete{
		manifestReader,
		options,
	}
}

func (d Delete) appName(baseName string, postfix string) string {
	return fmt.Sprintf("%s-%s", baseName, postfix)
}

func (d Delete) DeletePlan() (Plan, error) {
	application, err := d.manifestReader.Read(d.options.ManifestPath)
	if err != nil {
		return Plan{}, err
	}

	appName := d.appName(application.Name, d.options.Postfix)
	return Plan{[]Cmd{
		CfCmd{[]string{"delete", appName, "-f", "-r"}},
	}}, nil
}
