package issue

import (
	"github.com/sniperkit/github-release-checker/checker/vcs/github/repo"
)

type Issue struct {
	Repository *repo.Repository
	Version    string
}

func newIssue(repo *repo.Repository, version string) *Issue {
	return &Issue{
		Repository: repo,
		Version:    version,
	}
}
