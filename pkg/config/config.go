package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ApiInfo ApiInfo
	Thread  Thread
}

type ApiInfo struct {
	ApiKey      string `toml:"apiKey"`
	ApiPassword string `toml:"apiPassword"`
}

type Thread struct {
	ThreadNum int `toml:"threadNum"`
}

const CONFIG_FILE_PATH = "./config.toml"

func LoadConfig() (*Config, error) {
	config := new(Config)
	_, err := toml.DecodeFile(CONFIG_FILE_PATH, config)
	if err != nil {
		log.Println("config parse error.")
		return nil, err
	}

	return config, nil
}
