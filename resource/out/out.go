package out

import (
	. "github.com/simonjohansson/cf-protocol/command"
	"code.cloudfoundry.org/cli/cf/errors"
	"path"
	"reflect"
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

func (o Out) workingDir() string {
	appRoot := path.Join(o.sourceRoot, o.input.Params.Path)
	return appRoot
}

func (o Out) cfLogin() CliCmd {
	return CliCmd{
		Args: []string{"cf", "login",
			"-a", o.input.Source.Api,
			"-u", o.input.Source.Username,
			"-p", o.input.Source.Password,
			"-o", o.input.Params.Org,
			"-s", o.input.Params.Space,
		},
	}
}

func (o Out) protocolCommand() CliCmd {
	var cmd CliCmd
	commandName := "protocol-" + o.input.Params.Cmd
	switch commandName {
	case "protocol-push":
		cmd = CliCmd{
			Args: []string{
				"cf", commandName,
				"-manifest", o.input.Params.ManifestPath,
				"-domain", o.input.Params.TestDomain,
				"-postfix", o.concourseEnv.BuildName,
			},
			Dir: o.workingDir(),
		}
	case "protocol-promote", "protocol-cleanup", "protocol-delete":
		cmd = CliCmd{
			Args: []string{
				"cf", commandName,
				"-manifest", o.input.Params.ManifestPath,
				"-postfix", o.concourseEnv.BuildName,
			},
			Dir: o.workingDir(),
		}
	}
	return cmd
}

type MyStruct struct {
	A, B, C string
}

func (o Out) errorIfMissingSourceAndParamsValues() error {

	msValuePtr := reflect.ValueOf(&o.input.Params)
	msValue := msValuePtr.Elem()
	val := reflect.Indirect(reflect.ValueOf(o.input.Params))
	for i := 0; i < msValue.NumField(); i++ {
		value := msValue.Field(i).String()
		name := val.Type().Field(i).Name
		if value == "" {
			if o.input.Params.Cmd != "push" && o.input.Params.TestDomain == "" {
				// We only need testDomain if cmd is push
			} else {
				return errors.New("params." + name + " must be set!")
			}
		}
	}

	msValuePtr = reflect.ValueOf(&o.input.Source)
	msValue = msValuePtr.Elem()
	val = reflect.Indirect(reflect.ValueOf(o.input.Source))
	for i := 0; i < msValue.NumField(); i++ {
		value := msValue.Field(i).String()
		name := val.Type().Field(i).Name
		if value == "" {
			return errors.New("source." + name + " must be set!")
		}
	}

	return nil
}

func (o Out) OutPlan() (Plan, error) {
	if (o.input == &Input{}) {
		return Plan{}, errors.New("Input looks empty..")
	}
	err := o.errorIfMissingSourceAndParamsValues()
	if err != nil {
		return Plan{}, err
	}

	return Plan{
		[]Cmd{
			o.cfLogin(),
			o.protocolCommand(),
		},
	}, nil
}
