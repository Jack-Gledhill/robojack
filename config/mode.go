package config

import "os"

const (
	DevMode = "development"
	// ProdMode is the MODE value that will put the bot in production mode
	ProdMode = "production"
)

var mode = os.Getenv("MODE")

func IsDevelopment() bool {
	return !IsProduction()
}

func IsProduction() bool {
	return mode == ProdMode
}

func Mode() string {
	if IsProduction() {
		return ProdMode
	} else if IsDevelopment() {
		return DevMode
	} else {
		return "unknown"
	}
}
