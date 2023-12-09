package main

import (
	"log"
	"task-manager/config"
	"task-manager/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
	}

	app.Run(cfg)
}
