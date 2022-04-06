package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/horzu/golang/cart-api/internal/order"
	"github.com/horzu/golang/cart-api/internal/product"
	"github.com/horzu/golang/cart-api/pkg/config"
	db "github.com/horzu/golang/cart-api/pkg/database"
	"github.com/horzu/golang/cart-api/pkg/graceful"
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

	r:= gin.Default()

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.ServerConfig.Port),
		Handler: r,
		ReadTimeout: time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int(time.Second)),
	}

	rootRouter := r.Group(cfg.ServerConfig.RouterPrefix)

	orderRouter := rootRouter.Group("/orders")
	productRouter := rootRouter.Group("/products")


	// Order Repository
	orderRepo := order.NewOrderRepository(DB)
	orderRepo.Migration()
	order.NewOrderHandler(orderRouter, orderRepo)

	
	// Product Repository
	productRepo := product.NewProductRepository(DB)
	productRepo.Migration()
	product.NewProductHandler(productRouter, productRepo)

	go func(){
		if err:= srv.ListenAndServe(); err!=http.ErrServerClosed{
			log.Fatalf("listen error: %v", err)
		}
	}()

	log.Println("Shopping Cart service started!")

	graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int(time.Second)))
}