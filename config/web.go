package config

var Web *WebConfig

func init() {
	Web = &file.Web
}

type WebConfig struct {
	Port string `yaml:"port"`
}
