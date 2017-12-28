package language

import (
	"github.com/sniperkit/github-release-checker/checker/vcs/github/repo"
)

type Language struct {
	Repository *repo.Repository
	Version    string
}

func newLanguage(repo *repo.Repository, version string) *Language {
	return &Language{
		Repository: repo,
		Version:    version,
	}
}
