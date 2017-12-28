package handlers

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/sniperkit/github-release-checker/checker/vcs/github/repo/tag"
	"github.com/sniperkit/github-release-checker/helper/logging"
)

var (
	ErrHandlerAlreadyExists = errors.New("handler already exists")
)

var (
	handlers      = make(map[string]Handler)
	handlersMutex sync.RWMutex
)

type NewHandlerFunc func() (Handler, error)

func Handle(tagChan tag.Chan, done chan<- struct{}) error {
	handlersMutex.RLock()
	defer handlersMutex.RUnlock()

	copies := tag.CloneChan(tagChan, len(handlers))
	wg := &sync.WaitGroup{}
	defer func() {
		wg.Wait()
		close(done)
	}()

	i := 0
	for name, handler := range handlers {
		logging.Debugf("notifying handler '%s'", name)
		c := copies[i]
		wg.Add(1)
		go func(handler Handler) {
			defer wg.Done()
			handler.Handle(c)
		}(handler)
		i++
	}

	return nil
}

func Register(name string, f NewHandlerFunc) error {
	handlersMutex.Lock()
	defer handlersMutex.Unlock()
	if _, ok := handlers[name]; ok {
		return ErrHandlerAlreadyExists
	}
	handler, err := f()
	if err != nil {
		return errors.Wrap(err, "failed to create handler")
	}
	handlers[name] = handler
	return nil
}
