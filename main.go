package main

import (
	"fmt"
	"log"
	"kyros-logger/config"
	"kyros-logger/logger"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded config: %+v\n", cfg)
	logger, err := logger.NewLogger(cfg)
	if err != nil {
		log.Fatal(err)
	}
	logger.Infof("Application started\n")
	fmt.Println("Kyros Logger is running...")
}