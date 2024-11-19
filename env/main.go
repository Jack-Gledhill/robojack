package env

import "os"

const (
	// DevMode is the MODE value that will put the bot in development mode
	DevMode = "development"
	// ProdMode is the inverse of DevMode
	ProdMode = "production"
)

// Development will return true if the bot is in development mode
func Development() bool {
	m := os.Getenv("MODE")
	return m == DevMode || m == ""
}

// Production will return true if the bot is in production mode
func Production() bool {
	m := os.Getenv("MODE")
	return m == ProdMode
}

// Mode returns the current mode of the bot as a string
func Mode() string {
	if Development() {
		return DevMode
	} else if Production() {
		return ProdMode
	} else {
		return "unknown"
	}
}
