package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const FilePath = "config.yml"

var file FileConfig

type FileConfig struct {
	Bot    BotConfig   `yaml:"bot"`
	Emojis EmojiConfig `yaml:"emojis"`
}

func init() {
	f, err := os.ReadFile(FilePath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(f, &file)
	if err != nil {
		panic(err)
	}
}
