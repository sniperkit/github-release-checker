package user

import (
	"github.com/google/go-github/github"
)

type Organization struct {
	*github.Organization
	IsIgnored   bool
	IsImportant bool
}

func newOrganization(org *github.Organization) *Organization {
	return &Organization{
		Organization: org,
	}
}
