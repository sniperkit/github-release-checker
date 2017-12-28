package repo

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/sniperkit/github-release-checker/helper/logging"
)

func listStarredAll(activity *github.ActivityService, repoChan chanW, wg *sync.WaitGroup) error {
	firstPage := 1
	res, err := listStarred(activity, firstPage, repoChan)
	if err != nil {
		return errors.Wrap(err, "failed to get initial starred page")
	}
	for page := firstPage + 1; page <= res.LastPage; page++ {
		wg.Add(1)
		go func(page int) {
			defer func() {
				wg.Done()
				logging.Debug("done getting starred repos, page", page)
			}()
			logging.Debug("start getting starred repos, page", page)
			time.Sleep(200 * time.Millisecond)
			if _, err := listStarred(activity, page, repoChan); err != nil {
				logging.Error(errors.Wrap(err, fmt.Sprintf("failed to get starred repos, page %d", page)))
			}
		}(page)
	}
	return nil
}

func listStarred(activity *github.ActivityService, page int, repoChan chanW) (*github.Response, error) {
	opt := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}
	// TODO(leon): Pass in context
	time.Sleep(200 * time.Millisecond)
	starred, res, err := activity.ListStarred(context.Background(), "", opt)
	if err != nil {
		return nil, err
	}
	for _, star := range starred {
		repoChan <- newRepository(star.Repository)
	}
	return res, nil
}
