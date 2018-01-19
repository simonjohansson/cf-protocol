package main

import (
	"os"
	. "github.com/simonjohansson/cf-protocol/helpers"
	. "github.com/simonjohansson/cf-protocol/resource/out"
	"io/ioutil"
	"encoding/json"
	"github.com/caarlos0/env"
	"time"
	"github.com/simonjohansson/cf-protocol/command"
	"fmt"
)

func getData(data os.File) (Input, ConcourseEnv, error) {
	inbytes, err := ioutil.ReadAll(&data)
	var input Input
	err = json.Unmarshal(inbytes, &input)
	if err != nil {
		return Input{}, ConcourseEnv{}, err
	}
	var concourseEnv ConcourseEnv
	err = env.Parse(&concourseEnv)
	if err != nil {
		return Input{}, ConcourseEnv{}, err
	}
	return input, concourseEnv, nil
}

func logErrorAndExit(err error, logger Logger) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func main() {
	startTime := time.Now()

	logger := NewLogger()
	logger.ForwardStdoutToStderr() // stdout in a out resource must ONLY be used to return a result.

	sourceRoot := os.Args[1]
	input, concourseEnv, err := getData(*os.Stdin)
	logErrorAndExit(err, logger)

	plan, err := NewOut(sourceRoot, &input, &concourseEnv).OutPlan()
	logErrorAndExit(err, logger)

	logger.Info("Execution plan ")
	plan.PrintPlan(logger)
	logger.Info("Executing")
	err = command.NewCliExecutor(logger).Execute(plan)
	logErrorAndExit(err, logger)

	finishTime := time.Now()
	durationString := fmt.Sprintf("%s", finishTime.Sub(startTime).String())

	response := Response{
		Version: Version{
			Timestamp: finishTime,
		},
		Metadata: []MetadataPair{
			MetadataPair{Name: "Api", Value: input.Source.Api},
			MetadataPair{Name: "Org", Value: input.Params.Org},
			MetadataPair{Name: "Space", Value: input.Params.Space},
			MetadataPair{Name: "Duration", Value: durationString},
		},
	}
	json.NewEncoder(os.Stdout).Encode(response)
}
