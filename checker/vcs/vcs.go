package vcs

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/color"
	"github.com/hoop33/entrevista"
	// "github.com/sniperkit/github-release-checker/checker/model"
)

// Vcs represents a vcs service
type Vcs interface {
	Login(ctx context.Context) (string, error)
	// GetItem(ctx context.Context, token, entity, opts Opts)
	// GetItems(ctx context.Context, itemChan chan<- *model.ItemResult, token, entity, , opts Opts)
}

var vcs = make(map[string]Vcs)

func registerVcs(v Vcs) {
	vcs[Name(v)] = v
}

// Name returns the name of a service
func Name(v Vcs) string {
	parts := strings.Split(reflect.TypeOf(v).String(), ".")
	return strings.ToLower(parts[len(parts)-1])
}

// ForName returns the service for a given name, or an error if it doesn't exist
func ForName(name string) (Vcs, error) {
	if v, ok := vcs[strings.ToLower(name)]; ok {
		return v, nil
	}
	return &NotFound{}, fmt.Errorf("Vcs '%s' not found", name)
}

func createInterview() *entrevista.Interview {
	interview := entrevista.NewInterview()
	interview.ShowOutput = func(message string) {
		fmt.Print(color.GreenString(message))
	}
	interview.ShowError = func(message string) {
		color.Red(message)
	}
	return interview
}
