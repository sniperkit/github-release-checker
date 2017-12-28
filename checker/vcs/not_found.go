package vcs

import (
	"context"
	"errors"
	// "github.com/sniperkit/github-release-checker/checker/model"
)

// NotFound is used when the specified service is not found
type NotFound struct {
}

// Login is not implemented
func (nf *NotFound) Login(ctx context.Context) (string, error) {
	return "", errors.New("VCS Service not found")
}

// GetItem is not implemented
// func (nf *NotFound) GetItem(ctx context.Context, token string, entity string, opts Opts) {}

// GetItems is not implemented
// func (nf *NotFound) GetItems(ctx context.Context, itemChan chan<- *model.ItemResult, token string, entity string, opts Opts) {}
