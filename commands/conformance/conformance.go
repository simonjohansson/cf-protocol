package conformance

import (
	"encoding/json"
	"io/ioutil"
	. "github.com/simonjohansson/cf-protocol/helpers"
	"flag"
)

type Revision struct {
	Revision string `json:"revision"`
}

func (r Revision) asJsonString() string {
	json, _ := json.Marshal(r)
	return string(json)
}

func checkRevision(appUrl string, fetcher HttpClient, log Logger) error {

	versionUlr := appUrl + "/internal/version"
	response, err := fetcher.Get(versionUlr)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var revision Revision
	err = json.Unmarshal([]byte(body), &revision)
	if err != nil {
		return err
	}

	log.Info("/internal/version gives us: " + revision.asJsonString())
	return nil
}

func checkInternalStatus(appUrl string, fetcher HttpClient, logger Logger) error {
	logger.Info("Checking that app returns 200 on /internal/status")
	statusUrl := appUrl + "/internal/status"
	_, err := fetcher.Get(statusUrl)

	return err
}

func Conformance(appUrl string, fetcher HttpClient, logger Logger) error {
	logger.Info("Checking that app has /internal/version")
	err := checkRevision(appUrl, fetcher, logger)
	if err != nil {
		return err
	}

	err = checkInternalStatus(appUrl, fetcher, logger)
	if err != nil {
		return err
	}

	return nil
}

func parseArgs(logger Logger, flagset *flag.FlagSet, args []string) error {
	err := flagset.Parse(args[1:])
	if err != nil {
		return err
	}

	return nil
}

func RunConformance(logger Logger, args []string) error {

	flagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	appUrl := flagSet.String("appUrl", "", "app url to push app to, run confirmance aginst etc.")
	err := parseArgs(logger, flagSet, args)
	if err != nil {
		return err
	}

	logger.Info("Starting conformance on app with url '" + *appUrl + "'")
	httpClient := NewHttpClient()
	err = Conformance(*appUrl, httpClient, logger)
	if err != nil {
		return err
	}

	logger.Info("Conformance succeeded!")
	return nil
}
