package main

import (
	"code.cloudfoundry.org/cli/plugin"
	. "github.com/simonjohansson/cf-protocol/commands/conformance"
	. "github.com/simonjohansson/cf-protocol/commands/promote"
	. "github.com/simonjohansson/cf-protocol/commands/push"
	. "github.com/simonjohansson/cf-protocol/helpers"
	"syscall"
	. "github.com/simonjohansson/cf-protocol/commands/delete"
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
			{
				Name:     "protocol-push",
				HelpText: "Pushes the app",
				UsageDetails: plugin.Usage{
					Usage: "protocol-push -domain -postfix -manifest ",
				},
			},
			{
				Name:     "protocol-promote",
				HelpText: "Promotes the app",
				UsageDetails: plugin.Usage{
					Usage: "protocol-promote -manifest ",
				},
			},
			{
				Name:     "protocol-delete",
				HelpText: "Deletes the app",
				UsageDetails: plugin.Usage{
					Usage: "protocol-prodeletemote -manifest -postfix",
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

func (c *protocol) Run(cliConnection plugin.CliConnection, args []string) {
	logger := NewLogger()
	switch args[0] {
	case "protocol-conformance":
		err := RunConformance(logger, args)
		if err != nil {
			logger.Error("Conformance failed due to " + err.Error())
			syscall.Exit(-1)
		}
	case "protocol-push":
		err := RunPush(cliConnection, logger, args)
		if err != nil {
			logger.Error("Push failed due to " + err.Error())
			syscall.Exit(-1)
		}
	case "protocol-promote":
		err := RunPromote(cliConnection, logger, args)
		if err != nil {
			logger.Error("Push failed due to " + err.Error())
			syscall.Exit(-1)
		}
	case "protocol-delete":
		err := RunDelete(cliConnection, logger, args)
		if err != nil {
			logger.Error("Delete failed due to " + err.Error())
			syscall.Exit(-1)
		}
	}
}
