package github

import (
	"golang.org/x/oauth2"
	"net/http"

	// httpclient "github.com/sniperkit/github-release-checker/helper/http"
	"github.com/google/go-github/github"
)

type Github struct {
	config *Config
	httpc  *http.Client
	Client *github.Client
}

func New(c *Config) *Github {
	return &Github{
		config: c,
		Client: github.NewClient(&http.Client{
			Transport: &oauth2.Transport{
				Source: oauth2.StaticTokenSource(
					&oauth2.Token{AccessToken: c.Token},
				),
			},
		}),
	}
}

func NewWithTransport(c *Config, ht http.RoundTripper) *Github {
	return &Github{
		config: c,
		Client: github.NewClient(&http.Client{
			Transport: &oauth2.Transport{
				Base: ht,
				Source: oauth2.StaticTokenSource(
					&oauth2.Token{AccessToken: c.Token},
				),
			},
		}),
	}
}

func NewBasicAuth(c *Config) *Github {
	bat := github.BasicAuthTransport{
		Username: c.User,
		Password: c.Token,
	}
	return &Github{
		config: c,
		Client: github.NewClient(bat.Client()),
	}
}
