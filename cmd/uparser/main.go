package main

import (
	"flag"
	"fmt"
	"log"
	"nek-back/internal/uparser/config"
	"os"
)

func main() {
	// Init flags
	uip := flag.String("uip", "", "uip protocol key")
	configFlag := flag.String("config", "./config/config.dev.json", "configuration file")
	flag.Parse()

	// Get configuration
	cfg, err := config.Init(*uip, *configFlag)
	if err != nil {
		log.Fatal(err)
	}

	// Setup logger
	// TODO log file path
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		//log.Fatal(err)
	}

	log.SetOutput(file)

	// Run main app
	// app.Run(config)

	fmt.Println("Hello world! ", *cfg)
}
