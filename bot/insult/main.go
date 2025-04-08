package insult

import (
	"os"

	"gopkg.in/yaml.v3"
)

const MasterListPath = "insults.yml"

var MasterList Insults

type Insults struct {
	BG3         []string `yaml:"bg3"`
	DnD         []string `yaml:"dnd"`
	General     []string `yaml:"general"`
	Programming []string `yaml:"programming"`
	SCS         []string `yaml:"scs"`
	UoS         []string `yaml:"uos"`
}

func init() {
	f, err := os.ReadFile(MasterListPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(f, &MasterList)
	if err != nil {
		panic(err)
	}
}
