package tag

import (
	"github.com/sniperkit/github-release-checker/checker/vcs/github/repo"
)

type Tag struct {
	Repository *repo.Repository
	Version    string
}

func newTag(repo *repo.Repository, version string) *Tag {
	return &Tag{
		Repository: repo,
		Version:    version,
	}
}
