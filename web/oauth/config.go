package oauth

import (
	"github.com/Jack-Gledhill/robojack/config"

	"golang.org/x/oauth2"
)

var Config = &oauth2.Config{
	ClientID:     config.Web.OAuth.ClientID,
	ClientSecret: config.Web.OAuth.ClientSecret,
	Endpoint: oauth2.Endpoint{
		AuthURL:   "https://discord.com/api/oauth2/authorize",
		TokenURL:  "https://discord.com/api/oauth2/token",
		AuthStyle: oauth2.AuthStyleInParams,
	},
	RedirectURL: config.Web.Domain().String() + "/auth/callback",
	Scopes:      []string{"identify"},
}
