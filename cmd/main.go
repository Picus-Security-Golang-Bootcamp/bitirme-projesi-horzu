package main

import (
	"log"

	"github.com/horzu/golang/cart-api/pkg/config"
	db "github.com/horzu/golang/cart-api/pkg/database"
	logger "github.com/horzu/golang/cart-api/pkg/logging"
)

func main() {
	log.Println("Cart service starting...")

	// set env for local development
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err!=nil{
		log.Fatalf("loadconfig failed: %v", err)
	}

	// Set gloabal logger
	logger.NewLogger(cfg)
	defer logger.Close()

	DB := db.Connect(cfg)
	
}