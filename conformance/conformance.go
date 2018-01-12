package conformance

import (
	"encoding/json"
	"io/ioutil"
	"github.com/simonjohansson/cf-protocol/logger"
)

type Revision struct {
	Revision string `json:"revision"`
}

func (r Revision) asJsonString() string {
	json, _ := json.Marshal(r)
	return string(json)
}

func checkRevision(appUrl string, fetcher HttpFetcher, log logger.Logger) error {

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

func checkInternalStatus(appUrl string, fetcher HttpFetcher, logger logger.Logger) error {
	logger.Info("Checking that app returns 200 on /internal/status")
	statusUrl := appUrl + "/internal/status"
	_, err := fetcher.Get(statusUrl)

	return err
}

func Conformance(appUrl string, fetcher HttpFetcher, logger logger.Logger) error {
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
