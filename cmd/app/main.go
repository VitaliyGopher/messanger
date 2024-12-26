package main

import (
	"fmt"
	"log"

	"github.com/VitaliyGopher/messanger/internal/config"
	"github.com/VitaliyGopher/messanger/internal/router"
)

var (
	CONFIG_PATH string = "config/config.toml"
)

func main() {
	cfg, err := config.NewConfig(CONFIG_PATH)
	if err != nil {
		log.Fatalf("Error with config: %s", err)
	}

	r := router.NewRouter()

	addr := fmt.Sprintf("%s%s", cfg.Host, cfg.Port)
	r.Run(addr)
}

// TODO: Телефон -> код (generate + verefication)       
// 		 JWT generate token
//       Refreshing token    