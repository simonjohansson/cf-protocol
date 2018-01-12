package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"github.com/simonjohansson/cf-protocol/conformance"
	"syscall"
	"github.com/simonjohansson/cf-protocol/logger"
	"flag"
	"os"
)

type protocol struct{}

func (c *protocol) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "protocol",
		Commands: []plugin.Command{
			{
				Name:     "protocol-conformance",
				HelpText: "Makes a conformance check on app-url",
				UsageDetails: plugin.Usage{
					Usage: "protocol-conformance -appUrl",
				},
			},
		},
		Version: plugin.VersionType{
			0, 0, 1,
		},
	}
}

func main() {
	plugin.Start(new(protocol))
}
func ParseArgs(args []string) error {
	return nil
}

func (c *protocol) Run(cliConnection plugin.CliConnection, args []string) {
	logger := logger.NewLogger()
	if args[0] == "protocol-conformance" {
		runConformance(logger, args)
	}
}

func parseArgs(logger logger.Logger, flagset *flag.FlagSet, args []string) {
	err := flagset.Parse(args[1:])
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func runConformance(logger logger.Logger, args []string) {

	flagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	appUrl := flagSet.String("appUrl", "", "app url to push app to, run confirmance aginst etc.")
	parseArgs(logger, flagSet, args)

	logger.Info("Starting conformance on app with url '" + *appUrl + "'")
	httpClient := conformance.NewHttpClient()
	err := conformance.Conformance(*appUrl, httpClient, logger)
	if err != nil {
		logger.Error("Conformance failed due to " + err.Error())
		syscall.Exit(-1)
	}
	logger.Info("Conformance succeeded!")
}
