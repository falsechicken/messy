package messy

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server        string
	Username      string
	Password      string
	Remote        string
	Status        string
	StatusMessage string
	NoTLS         bool
	StartTLS      bool
	DebugMode     bool
	Session       bool
}

func CreateDefaultConfig() bool {
	return true
}

func ReadConfig(location string) Config {
	_, err := os.Stat(location)
	if err != nil {
		log.Fatal("Config file is missing: ", location)
	}

	var config Config
	if _, err := toml.DecodeFile(location, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}
