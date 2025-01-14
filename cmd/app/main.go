package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/VitaliyGopher/messanger/internal/app/server"
)

var (
	CONFIG_PATH string = "config/config.toml"
)

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

// TODO: Телефон -> код (generate + verefication)
// 		 JWT generate token
//       Refreshing token
