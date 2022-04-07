package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/horzu/golang/cart-api/internal/auth"
	"github.com/horzu/golang/cart-api/internal/order"
	"github.com/horzu/golang/cart-api/internal/product"
	"github.com/horzu/golang/cart-api/pkg/config"
	db "github.com/horzu/golang/cart-api/pkg/database"
	"github.com/horzu/golang/cart-api/pkg/graceful"
	logger "github.com/horzu/golang/cart-api/pkg/logging"
	"go.uber.org/zap"
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

	gin.SetMode(gin.ReleaseMode)
	r:= gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string{
		// custom format arranged for logger
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
	}))

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.ServerConfig.Port),
		Handler: r,
		ReadTimeout: time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int(time.Second)),
	}

	rootRouter := r.Group(cfg.ServerConfig.RouterPrefix)

	orderRouter := rootRouter.Group("/orders")
	productRouter := rootRouter.Group("/products")
	authRouter := rootRouter.Group("/user")


	// Order Repository
	orderRepo := order.NewOrderRepository(DB)
	orderRepo.Migration()
	order.NewOrderHandler(orderRouter, orderRepo)

	
	// Product Repository
	productRepo := product.NewProductRepository(DB)
	productRepo.Migration()
	product.NewProductHandler(productRouter, productRepo)

	auth.NewAuthHandler(authRouter, cfg)

	go func(){
		if err:= srv.ListenAndServe(); err!=http.ErrServerClosed{
			log.Fatalf("listen error: %v", err)
		}
	}()

	r.GET("healthx", func(c *gin.Context){
		c.JSON(http.StatusOK, nil)
	})

	r.GET("readyx", func (c *gin.Context){
		db, err := DB.DB()
		if err!=nil{
			zap.L().Fatal("Cannot get sql database instance ", zap.Error(err))
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		if err := db.Ping(); err!=nil{
			zap.L().Fatal("Cannot ping database ", zap.Error(err))
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	log.Println("Shopping Cart service started!")

	graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int(time.Second)))
}