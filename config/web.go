package config

import "net/url"

var Web *WebConfig

func init() {
	Web = &file.Web
}

type WebConfig struct {
	FullDomain string      `yaml:"domain"`
	JWT        JWTConfig   `yaml:"jwt"`
	OAuth      OAuthConfig `yaml:"oauth"`
	Port       string      `yaml:"port"`
}

func (w *WebConfig) Domain() *url.URL {
	u, err := url.Parse(w.FullDomain)
	if err != nil {
		panic(err)
	}

	return u
}

type JWTConfig struct {
	SigningSecret string `yaml:"signing_secret"`
	Validity      int    `yaml:"validity"`
}

type OAuthConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}
