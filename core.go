package app

import (
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
}

// ================================================================
//
// ================================================================
func New() (*App, error) {
	loc, err := time.LoadLocation(os.Getenv("TIMEZONE"))
	if err != nil {
		return nil, err
	}

	env := &App{
		AppTitle:   os.Getenv("APP_TITLE"),
		AppHost:    os.Getenv("APP_HOST"),
		AppPath:    path.Join("/", os.Getenv("APP_PATH")),
		AppPort:    os.Getenv("APP_PORT"),
		GinMode:    os.Getenv("GIN_MODE"),
		Location:   loc,
		TrustProxy: os.Getenv("TRUST_PROXY"),
	}

	env.AppRootUrl, err = url.ParseRequestURI("https://" + path.Join(env.AppHost, env.AppPath))
	return env, err
}
