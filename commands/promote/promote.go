package promote

import (
	"code.cloudfoundry.org/cli/plugin"
	. "github.com/simonjohansson/cf-protocol/command"
	. "github.com/simonjohansson/cf-protocol/helpers"
)

type Promote struct {
	cliConnection plugin.CliConnection
	options       Options
}

func NewPromote(cliConnection plugin.CliConnection, options Options) Promote {
	return Promote{
		cliConnection,
		options,
	}
}

func (p Promote) PromotePlan() (Plan, error) {
	return Plan{
		Cmds: []Cmd{
			CfApps{},
		},
	}, nil
}
