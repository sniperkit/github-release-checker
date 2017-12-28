package readme

import (
	"github.com/sniperkit/github-release-checker/checker/vcs/github/repo"
)

type Readme struct {
	Repository *repo.Repository
	Version    string
}

func newReadme(repo *repo.Repository, version string) *Readme {
	return &Readme{
		Repository: repo,
		Version:    version,
	}
}
