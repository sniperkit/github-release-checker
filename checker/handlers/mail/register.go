package mail

import (
	"crypto/tls"

	"github.com/sniperkit/github-release-checker/checker/handlers"
)

const (
	handlerName = "mail"
)

func Register(c *Config) error {
	if !c.Enabled {
		return nil
	}
	c.TLSConfig = &tls.Config{InsecureSkipVerify: c.Insecure}
	return handlers.Register(handlerName, func() (handlers.Handler, error) {
		return &mailHandler{config: c}, nil
	})
}
