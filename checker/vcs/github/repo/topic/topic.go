package topic

import (
	"github.com/sniperkit/github-release-checker/checker/vcs/github/repo"
)

type Topic struct {
	Repository *repo.Repository
	Version    string
}

func newTopic(repo *repo.Repository, version string) *Topic {
	return &Topic{
		Repository: repo,
		Version:    version,
	}
}
