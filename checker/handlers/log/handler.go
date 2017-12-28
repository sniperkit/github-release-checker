package log

import (
	"github.com/sniperkit/github-release-checker/checker/vcs/github/repo/tag"

	"github.com/sniperkit/github-release-checker/helper/logging"
)

type logHandler struct {
}

func (lh *logHandler) Handle(tagChan tag.Chan) {
	for tag := range tagChan {
		name := tag.Repository.GetFullName()
		version := tag.Version
		msg := "found new tag"
		if tag.Repository.IsImportant {
			msg = "found new important tag"
		}
		logging.Info(msg, name, version)
	}
}
