package user

import (
	"github.com/google/go-github/github"
)

type User struct {
	*github.User
	IsIgnored   bool
	IsImportant bool
}

func newUser(user *github.User) *User {
	return &User{
		User: user,
	}
}
