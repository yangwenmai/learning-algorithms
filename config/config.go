package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

const (
	configTOML = "config.toml"
)

// Config config
type Config struct {
	Username string
	Password string
}

func GetConfig() *Config {
	cfg := new(Config)
	if _, err := toml.DecodeFile(configTOML, &cfg); err != nil {
		log.Panicf(err.Error())
	}
	// log.Printf("get config: %s", cfg)
	return cfg
}
