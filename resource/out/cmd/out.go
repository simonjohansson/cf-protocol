package main

import (
	"os"
	. "github.com/simonjohansson/cf-protocol/helpers"
	. "github.com/simonjohansson/cf-protocol/resource/out"
	"io/ioutil"
	"encoding/json"
	"github.com/caarlos0/env"
	"github.com/simonjohansson/cf-protocol/command"
)

func getData() (Input, ConcourseEnv, error) {
	inbytes, err := ioutil.ReadAll(os.Stdin)
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

func main() {
	logger := NewLogger()
	sourceRoot := os.Args[1]
	input, concourseEnv, err := getData()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	plan, err := NewOut(sourceRoot, &input, &concourseEnv).OutPlan()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("Execution plan ")
	plan.PrintPlan(logger)
	logger.Info("Executing")
	command.NewCliExecutor(logger).Execute(plan)
}
