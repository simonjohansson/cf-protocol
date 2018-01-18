package out

import "time"

type Params struct {
	Path         string `json:"path"`
	Cmd          string `json:"cmd"`
	Org          string `json:"org"`
	Space        string `json:"space"`
	ManifestPath string `json:"manifestPath"`
	TestDomain   string `json:"testDomain"`
}
type Source struct {
	Api      string `json:"api"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Input struct {
	Source Source `json:"source"`
	Params Params `json:"params"`
}

type ConcourseEnv struct {
	BuildId           string `env:"BUILD_ID"`
	BuildName         string `env:"BUILD_NAME"`
	BuildJobName      string `env:"BUILD_JOB_NAME"`
	BuildPipelineName string `env:"BUILD_PIPELINE_NAME"`
	AtcExternalUrl    string `env:"ATC_EXTERNAL_URL"`
	BuildTeamName     string `env:"BUILD_TEAM_NAME"`
}

type Response struct {
	Version  Version        `json:"version"`
	Metadata []MetadataPair `json:"metadata"`
}
type Version struct {
	Timestamp time.Time `json:"timestamp"`
}

type MetadataPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}