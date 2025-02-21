package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/VitaliyGopher/messanger/internal/app/server"
	"github.com/VitaliyGopher/messanger/pkg/env"
)

var (
	CONFIG_PATH string = "config/config.toml"
)

func init() {
	if err := env.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	config := server.NewConfig()
	_, err := toml.DecodeFile(CONFIG_PATH, &config)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(config); err != nil {
		log.Fatal(err)
	}
}
