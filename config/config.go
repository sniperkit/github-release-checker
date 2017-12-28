package config

import (
	"github.com/sniperkit/github-release-checker/checker"

	mailhandler "github.com/sniperkit/github-release-checker/checker/handlers/mail"
)

type Config struct {
	CheckerConfig *checker.Config     `toml:"checker"`
	MailConfig    *mailhandler.Config `toml:"mail"`
}
