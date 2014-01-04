package main

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/stvp/go-toml-config"
)

var (
	// GoogleClientID is the client ID for OAuthConfig.
	// It's stored inside the config TOML file for security reasons.
	GoogleClientID     = config.String("taskqueue.clientID", "Unknown")
	// GoogleClientSecret is the client secret for OAuthConfig.
	// It's stored inside the config TOML file for security reasons.
	GoogleClientSecret = config.String("taskqueue.clientSecret", "Unknown")
	// OAuthConfig is the used configuration for every Google API access.
	OAuthConfig        = &oauth.Config{
		Scope:    "https://www.googleapis.com/auth/taskqueue",
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://accounts.google.com/o/oauth2/token",
		// Use "postmessage" for the code-flow for server side apps
		RedirectURL: "postmessage",
	}
)

// ReadConfig parses a given TOML file and fills variables with the given information (mostly API secrets).
func ReadConfig() {
	config.Parse("../src/socialvibes/socialvibes.toml")
	OAuthConfig.ClientId = *GoogleClientID
	OAuthConfig.ClientSecret = *GoogleClientSecret
}