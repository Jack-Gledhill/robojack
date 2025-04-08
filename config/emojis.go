package config

var Emojis *EmojiConfig

func init() {
	Emojis = &file.Emojis
}

type EmojiConfig struct {
	Angry        string `yaml:"angry"`
	Crying       string `yaml:"crying"`
	Dead         string `yaml:"dead"`
	Facepalm     string `yaml:"facepalm"`
	MiddleFinger string `yaml:"middle_finger"`
	LaserEyes    string `yaml:"laser_eyes"`
	Love         string `yaml:"love"`
	Peace        string `yaml:"peace"`
	Peeking      string `yaml:"peeking"`
	Smiling      string `yaml:"smiling"`
	Sunglasses   string `yaml:"sunglasses"`
	Thinking     string `yaml:"thinking"`
	Waving       string `yaml:"waving"`
	Weary        string `yaml:"weary"`
	Wink         string `yaml:"wink"`
}
