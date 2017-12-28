package checker

import (
	"time"

	httpclient "github.com/sniperkit/github-release-checker/helper/http"

	"github.com/sniperkit/github-release-checker/checker/vcs/github"
	// "github.com/sniperkit/github-release-checker/checker/model"
)

type Config struct {
	CheckInterval duration `json:"interval" yaml:"interval" toml:"interval"`
	RequestDelay  duration `json:"request_delay" yaml:"request_delay" toml:"request_delay"`

	TransportConfig    *httpclient.Config  `json:"http" yaml:"http" toml:"http"`
	DBConfig           *DBConfig           `json:"db" yaml:"db" toml:"db"`
	RepositoriesConfig *RepositoriesConfig `json:"repositories" yaml:"repositories" toml:"repositories"`
	GithubConfig       *github.Config      `json:"github" yaml:"github" toml:"github"`
	// BitbucketConfig       *github.Config      `json:"bitbucket" yaml:"bitbucket" toml:"bitbucket"`
	// GitlabConfig       *github.Config      `json:"gitlab" yaml:"gitlab" toml:"gitlab"`
	// ServiceConfig      *github.Config      `json:"services" yaml:"services" toml:"services"`
}

type DBConfig struct {
	Path string `json:"path" yam;:"path" toml:"path"`
}

type RepositoriesConfig struct {
	Ignored   []string `json:"ignored" yaml:"ignored" toml:"ignored"`
	Important []string `json:"important" yaml:"important" toml:"important"`
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func (d duration) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}
