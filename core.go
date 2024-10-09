package app

import (
	"errors"
	"net/url"
	"os"
	"path"
	"time"
)

// ================================================================
//
// ================================================================
type App struct {
	AppRootUrl *url.URL
	AppTitle   string
	AppHost    string
	AppPath    string
	AppPort    string
	GinMode    string
	Location   *time.Location
	TrustProxy string
	Visibility string
}

// ================================================================
//
// ================================================================
func New() (*App, error) {
	loc, err := time.LoadLocation(os.Getenv("TIMEZONE"))
	if err != nil {
		return nil, err
	}

	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "debug" && ginMode != "release" && ginMode != "test" {
		return nil, errors.New("Invalid GIN_MODE value. (debug | release | test)")
	}

	visibility := os.Getenv("VISIBILITY")
	if visibility != "internal" && visibility != "external" {
		return nil, errors.New("Invalid VISIBILITY value. (internal | external)")
	}

	env := &App{
		AppTitle:   os.Getenv("APP_TITLE"),
		AppHost:    os.Getenv("APP_HOST"),
		AppPath:    path.Join("/", os.Getenv("APP_PATH")),
		AppPort:    os.Getenv("APP_PORT"),
		GinMode:    ginMode,
		Location:   loc,
		TrustProxy: os.Getenv("TRUST_PROXY"),
		Visibility: visibility,
	}

	env.AppRootUrl, err = url.ParseRequestURI("https://" + path.Join(env.AppHost, env.AppPath))
	return env, err
}
