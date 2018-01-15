package out

import (
	. "github.com/simonjohansson/cf-protocol/command"
	"code.cloudfoundry.org/cli/cf/errors"
	"path"
)

type Out struct {
	sourceRoot   string
	input        *Input
	concourseEnv *ConcourseEnv
}

func NewOut(sourceRoot string, input *Input, concourseEnv *ConcourseEnv) Out {
	return Out{
		sourceRoot:   sourceRoot,
		input:        input,
		concourseEnv: concourseEnv,
	}
}

func (o Out) changeWorkingDir() CliCmd {
	appRoot := path.Join(o.sourceRoot, o.input.Params.Dir)
	return CliCmd{
		[]string{"cd", appRoot},
	}
}

func (o Out) cfLogin() CliCmd {
	return CliCmd{
		[]string{"cf", "login",
			"-a", o.input.Source.Api,
			"-u", o.input.Source.Username,
			"-p", o.input.Source.Password,
			"-o", o.input.Source.Org,
			"-s", o.input.Source.Password,
		},
	}
}

func (o Out) protocolCommand() CliCmd {
	commandName := "protocol-" + o.input.Params.Cmd
	return CliCmd{
		[]string{
			"cf", commandName,
			"-manifest", o.input.Params.ManifestPath,
			"-domain", "domain.io",
			"-postfix", o.concourseEnv.BuildName,
		},
	}
}

func (o Out) errorIfMissingSourceAndParams() error {
	if o.input.Params.Dir == "" {
		return errors.New("params.dir must be set!")
	}
	return nil
}

func (o Out) OutPlan() (Plan, error) {
	if (o.input == &Input{}) {
		return Plan{}, errors.New("Input looks empty..")
	}
	err := o.errorIfMissingSourceAndParams()
	if err != nil {
		return Plan{}, err
	}
	return Plan{
		[]Cmd{
			o.changeWorkingDir(),
			o.cfLogin(),
			o.protocolCommand(),
		},
	}, nil
}
