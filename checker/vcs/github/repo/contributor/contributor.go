package contributor

import (
	"github.com/sniperkit/github-release-checker/checker/vcs/github/repo"
)

type Contributor struct {
	Repository *repo.Repository
	Version    string
}

func newContributor(repo *repo.Repository, version string) *Contributor {
	return &Contributor{
		Repository: repo,
		Version:    version,
	}
}
