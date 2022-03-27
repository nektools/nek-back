package app

import (
	"fmt"
	"log"
	"nek-back/internal/uparser/config"
	"os"
)

func Run(cfg *config.Config) {
	// Setup logger
	// TODO log file path
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		//log.Fatal(err)
	}
	log.SetOutput(file)

	fmt.Println(cfg)
}
