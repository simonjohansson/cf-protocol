package promote

import (
	"code.cloudfoundry.org/cli/plugin"
	. "github.com/simonjohansson/cf-protocol/command"
	. "github.com/simonjohansson/cf-protocol/helpers"
	"strings"
	"code.cloudfoundry.org/cli/util/manifest"
	"regexp"
	"fmt"
	"strconv"
	"code.cloudfoundry.org/cli/cf/errors"
)

type Promote struct {
	cliConnection  plugin.CliConnection
	manifestReader ManifestReader
	options        Options
}

func NewPromote(cliConnection plugin.CliConnection, manifestReader ManifestReader, options Options) Promote {
	return Promote{
		cliConnection,
		manifestReader,
		options,
	}
}

func (p Promote) createRoutesCmd(application manifest.Application) []Cmd {
	cmds := []Cmd{}
	appName := application.Name + "-" + p.options.Postfix
	for _, route := range application.Routes {
		parts := strings.Split(route, ".")
		hostname := parts[0]
		domain := strings.Join(parts[1:], ".")
		cmd := CfCmd{
			[]string{"map-route", appName, domain, "--hostname", hostname},
		}
		cmds = append(cmds, cmd)
	}
	return cmds
}

func (p Promote) getPostfixVersion(appName string) string {
	r, _ := regexp.Compile("^[a-zA-Z-]+-([0-9]+)$")
	return r.FindStringSubmatch(appName)[1]

}

func (p Promote) looksLikeSameApp(appName string, otherAppName string) bool {
	if otherAppName == fmt.Sprintf("%s-%s", appName, p.options.Postfix) {
		return false
	}

	r, _ := regexp.Compile(fmt.Sprintf("^%s-[0-9]+$", appName))
	return r.MatchString(otherAppName)
}

func (p Promote) createStopCmd(application manifest.Application) ([]Cmd, error) {
	apps, err := p.cliConnection.GetApps()
	if err != nil {
		return []Cmd{}, err
	}

	returnCmds := []Cmd{}
	for _, app := range apps {
		if p.looksLikeSameApp(application.Name, app.Name) {
			otherPostfix, _ := strconv.ParseInt(p.getPostfixVersion(app.Name), 10, 64)
			currentPostfix, _ := strconv.ParseInt(p.options.Postfix, 10, 64)
			if (otherPostfix > currentPostfix) {
				return []Cmd{}, errors.New(app.Name + " have a postfix greater than " + p.options.Postfix)
			}
			if app.State == "started" {
				cmd := CfCmd{
					[]string{"stop", app.Name},
				}
				returnCmds = append(returnCmds, cmd)
			}
		}
	}

	return returnCmds, nil
}

func (p Promote) PromotePlan() (Plan, error) {
	application, err := p.manifestReader.Read(p.options.ManifestPath)
	if err != nil {
		return Plan{}, err
	}

	routes := p.createRoutesCmd(application)
	stops, err := p.createStopCmd(application)
	if err != nil {
		return Plan{}, err
	}
	return Plan{
		Cmds: append(routes, stops...),
	}, nil
}
