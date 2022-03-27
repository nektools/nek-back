package main

import (
	"flag"
	"log"
	"nek-back/internal/uparser/app"
	"nek-back/internal/uparser/config"
)

func main() {
	// Init flags
	uip := flag.String("uip", "", "uip protocol key")
	configPath := flag.String("config", "./config/config.dev.json", "configuration file")
	flag.Parse()

	// Get configuration
	cfg, err := config.Init(uip, configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Run main app
	app.Run(cfg)
}
