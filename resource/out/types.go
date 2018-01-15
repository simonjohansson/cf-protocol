package out

type Params struct {
	Dir          string `json:"dir"`
	Cmd          string `json:"cmd"`
	ManifestPath string `json:"manifestPath"`
}
type Source struct {
	Api      string `json:"api"`
	Username string `json:"username"`
	Password string `json:"password"`
	Org      string `json:"org"`
	Space    string `json:"space"`
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