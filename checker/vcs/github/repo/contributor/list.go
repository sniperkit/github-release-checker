package contributor

import (
	"fmt"
	"strings"
	"sync"

	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"github.com/sniperkit/github-release-checker/checker/vcs/github/repo"
	"github.com/sniperkit/github-release-checker/helper/logging"
)

const (
	// TODO(leon): Make this configurable
	maxWorkers = 100
)

func newListWorker(repoChan repo.Chan, contributorChan chanW) {
	for repo := range repoChan {
		htmlURL := repo.GetHTMLURL()
		tagsURL := htmlURL + "/tags.atom"
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(tagsURL)
		if err != nil {
			logging.Error(errors.Wrap(err, fmt.Sprintf("failed to parse URL %s", tagsURL)))
			continue
		}
		for _, item := range feed.Items {
			split := strings.Split(item.Link, "/")
			version := split[len(split)-1]
			contributorChan <- newTag(repo, version)
		}
	}
}

func List(repoChan repo.Chan) Chan {
	contributorChan := make(chanRW)

	go func() {

		wg := &sync.WaitGroup{}
		defer func() {
			wg.Wait()
			close(contributorChan)
			logging.Debug("done listing tags")
		}()

		for i := 0; i < maxWorkers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				newListWorker(repoChan, onlyWritable(contributorChan))
			}()
		}
	}()

	return onlyReadable(tagChan)
}
