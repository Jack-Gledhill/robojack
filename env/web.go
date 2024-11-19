package env

import "os"

// WebPrefix is prepended to all web-related environment variables
const WebPrefix = "WEB_"

// Web wraps the environment variables related to the web server
var Web = WebConfig{
	Interface: DefaultVar(WebPrefix+"INTERFACE", "0.0.0.0"),
	Port:      DefaultVar(WebPrefix+"PORT", "3000"),
	Presence: PresenceConfig{
		DataDirectory: DefaultVar(WebPrefix+"PRESENCE_DATA_DIRECTORY", "/data/"),
	},
}

// WebConfig is a wrapper for Web
type WebConfig struct {
	Interface string
	Port      string
	Presence  PresenceConfig
}

// URL returns the full URL of the web server, using an external one if provided, or forming it from the interface and port otherwise
func (w *WebConfig) URL() string {
	u, ok := os.LookupEnv(WebPrefix + "URL")
	if !ok {
		return "http://" + w.ListenAddress()
	}

	return u
}

// ListenAddress creates an address that Gin can use to start the web server on
func (w *WebConfig) ListenAddress() string {
	return w.Interface + ":" + w.Port
}

// PresenceConfig holds config options about the presence API
type PresenceConfig struct {
	DataDirectory string
}
