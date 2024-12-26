package config

import "github.com/BurntSushi/toml"

type Config struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

func NewConfig(configPath string) (*Config, error) {
	var config *Config

	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}