package handlers

import (
	"github.com/sniperkit/github-release-checker/checker/vcs/github/repo/tag"
)

type Handler interface {
	Handle(tag.Chan)
}
