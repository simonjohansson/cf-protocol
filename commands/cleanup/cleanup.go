package cleanup

import (
	"fmt"
	. "github.com/simonjohansson/cf-protocol/helpers"
	. "github.com/simonjohansson/cf-protocol/command"
	. "code.cloudfoundry.org/cli/plugin"
	. "code.cloudfoundry.org/cli/plugin/models"
	"regexp"
	"sort"
)

type Cleanup struct {
	cliConnection  CliConnection
	manifestReader ManifestReader
	options        Options
}

func NewCleanup(cliConnection CliConnection, manifestReader ManifestReader, options Options) Cleanup {
	return Cleanup{
		cliConnection,
		manifestReader,
		options,
	}
}

type AppsSorter []GetAppsModel

func (a AppsSorter) Len() int           { return len(a) }
func (a AppsSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AppsSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

func (c Cleanup) appName(baseName string, postfix string) string {
	return fmt.Sprintf("%s-%s", baseName, postfix)
}

func (c Cleanup) stopCommands(apps []GetAppsModel) []Cmd {
	if len(apps) == 0 {
		return []Cmd{}
	}
	if(apps[0].State != "started") {
		return []Cmd{}
	}
	return []Cmd{CfCmd{[]string{"stop", apps[0].Name}}}
}

func (c Cleanup) deleteCommands(apps []GetAppsModel) []Cmd {
	if len(apps) <= 1 {
		return []Cmd{}
	}

	deleteCmds := []Cmd{}
	for _, app := range (apps[1:]) {
		cmd := CfCmd{[]string{"delete", app.Name, "-f"}}
		deleteCmds = append(deleteCmds, cmd)
	}
	return deleteCmds
}

func (c Cleanup) looksLikeSameApp(appName string, otherAppName string) bool {
	if otherAppName == fmt.Sprintf("%s-%s", appName, c.options.Postfix) {
		return false
	}

	r, _ := regexp.Compile(fmt.Sprintf("^%s-[0-9]+$", appName))
	return r.MatchString(otherAppName)
}

func (c Cleanup) interestingApps(appName string) ([]GetAppsModel, error) {
	apps, err := c.cliConnection.GetApps()
	if err != nil {
		return []GetAppsModel{}, err
	}

	interestingApps := []GetAppsModel{}
	for _, app := range apps {
		if c.looksLikeSameApp(appName, app.Name) {
			interestingApps = append(interestingApps, app)
		}
	}
	sort.Sort(sort.Reverse(AppsSorter(interestingApps)))
	return interestingApps, nil
}

func (c Cleanup) CleanupPlan() (Plan, error) {
	application, err := c.manifestReader.Read(c.options.ManifestPath)
	if err != nil {
		return Plan{}, err
	}

	appName := application.Name
	interestingApps, err := c.interestingApps(appName)
	if err != nil {
		return Plan{}, err
	}

	stopCommands := c.stopCommands(interestingApps)
	deleteCommands := c.deleteCommands(interestingApps)
	commands := append(stopCommands, deleteCommands...)
	return Plan{commands}, nil
}
