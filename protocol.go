package main

import (
	"code.cloudfoundry.org/cli/plugin"
	. "github.com/simonjohansson/cf-protocol/commands/promote"
	. "github.com/simonjohansson/cf-protocol/helpers"
	"syscall"
	. "github.com/simonjohansson/cf-protocol/commands/delete"
	. "github.com/simonjohansson/cf-protocol/commands/push"
	"github.com/simonjohansson/cf-protocol/command"
)

type protocol struct{}

func (c *protocol) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "protocol",
		Commands: []plugin.Command{
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

func executePlan(planName string, plan command.Plan, err error, logger Logger, cliConnection plugin.CliConnection) {
	if err != nil {
		logger.Error(planName + " failed due to " + err.Error())
		syscall.Exit(-1)
	}
	logger.Info("Running plan " + planName)
	logger.Info("Execution plan ")
	plan.PrintPlan(logger)
	logger.Info("Executing")
	err = plan.ExecutePlan(cliConnection, logger)
	if err != nil {
		logger.Error(err.Error())
		logger.Error("Aborting.")
		syscall.Exit(-1)
	}
	logger.Info("Finished!")
}

func (c *protocol) Run(cliConnection plugin.CliConnection, args []string) {
	logger := NewLogger()
	options, err := ParseArgs(args)
	if err != nil {
		logger.Error(err.Error())
		syscall.Exit(-1)
	}
	switch args[0] {
	case "protocol-push":
		plan, err := NewPush(NewManifestReader(), options).PushPlan()
		executePlan("Push", plan, err, logger, cliConnection)
	case "protocol-promote":
		plan, err := NewPromote(cliConnection, options).PromotePlan()
		executePlan("Push", plan, err, logger, cliConnection)
	case "protocol-delete":
		plan, err := NewDelete(NewManifestReader(), options).DeletePlan()
		executePlan("Push", plan, err, logger, cliConnection)
	}
}
