package commit

import (
	"github.com/sniperkit/github-release-checker/checker/vcs/github/repo"
)

type Commit struct {
	Repository *repo.Repository
	Version    string
}

func newCommit(repo *repo.Repository, version string) *Commit {
	return &Commit{
		Repository: repo,
		Version:    version,
	}
}
