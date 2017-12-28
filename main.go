package main

import (
	"flag"
	// "net/http"

	// notify
	logHandler "github.com/sniperkit/github-release-checker/checker/handlers/log"
	mailHandler "github.com/sniperkit/github-release-checker/checker/handlers/mail"

	// app
	"github.com/sniperkit/github-release-checker/checker"
	"github.com/sniperkit/github-release-checker/config"

	// persistent store
	_ "github.com/mattn/go-sqlite3"
	// pivot
	// gorm

	// helpers
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"github.com/sniperkit/github-release-checker/helper/logging"
	"github.com/sniperkit/github-release-checker/utils/pathutil"
)

const (
	defaultConfigFilePath = "$HOME/.github-release-checker/config.toml"
)

var (
	configFilePath = flag.String("config", defaultConfigFilePath, "optional, path where to find the config file")
	enableDebug    = flag.Bool("debug", false, "optional, whether to enable debug mode")
)

func main() {
	flag.Parse()

	if *enableDebug {
		logging.SetDebug()
	}

	if err := boot(); err != nil {
		logging.Fatal(err)
	}
}

func boot() error {
	cfp, err := annotateConfigFilePath(*configFilePath)
	if err != nil {
		return errors.Wrap(err, "failed to annotate config file path")
	}

	var conf config.Config
	if _, err := toml.DecodeFile(cfp, &conf); err != nil {
		return errors.Wrap(err, "failed to load or parse config file")
	}

	if err := logHandler.Register(); err != nil {
		logging.Error(errors.Wrap(err, "failed to register log handler"))
	}
	if err := mailHandler.Register(conf.MailConfig); err != nil {
		logging.Error(errors.Wrap(err, "failed to register mail handler"))
	}

	c, err := checker.New(conf.CheckerConfig)
	if err != nil {
		return errors.Wrap(err, "failed to create checker")
	}

	return c.Start()
}

func annotateConfigFilePath(p string) (string, error) {
	return pathutil.ReplaceHome(p)
}
