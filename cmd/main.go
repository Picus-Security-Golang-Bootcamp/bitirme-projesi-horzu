package main

import (
	"fmt"
	"log"

	"github.com/horzu/golang/cart-api/pkg/config"
)

func main() {
	log.Println("Cart service starting...")

	// set env for local development
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err!=nil{
		log.Fatalf("loadconfig failed: %v", err)
	}

	fmt.Println(cfg)
}