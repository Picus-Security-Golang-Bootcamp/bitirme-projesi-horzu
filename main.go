package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/horzu/golang/cart-api/internal/domain/cart"
	"github.com/horzu/golang/cart-api/internal/domain/cart/cartItem"
	"github.com/horzu/golang/cart-api/internal/domain/category"
	"github.com/horzu/golang/cart-api/internal/domain/order"
	"github.com/horzu/golang/cart-api/internal/domain/order/orderItem"
	"github.com/horzu/golang/cart-api/internal/domain/product"
	"github.com/horzu/golang/cart-api/internal/domain/users"
	"github.com/horzu/golang/cart-api/internal/domain/users/role"
	"github.com/horzu/golang/cart-api/pkg/config"
	db "github.com/horzu/golang/cart-api/pkg/database"
	"github.com/horzu/golang/cart-api/pkg/graceful"
	logger "github.com/horzu/golang/cart-api/pkg/logging"
	"go.uber.org/zap"
)

func main() {
	log.Println("Cart service starting...")

	// Set env for local development
	cfg, err := config.LoadConfig("./pkg/config/config-local")
	if err != nil {
		log.Fatalf("loadconfig failed: %v", err)
	}

	// Set gloabal logger
	logger.NewLogger(cfg)
	defer logger.Close()

	// Connect to database
	DB := db.Connect(cfg)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
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
		Addr:         fmt.Sprintf("127.0.0.1:%d", cfg.ServerConfig.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs * int(time.Second)),
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs * int(time.Second)),
	}

	rootRouter := r.Group(cfg.ServerConfig.RouterPrefix)

	// Routes
	authRouter := rootRouter.Group("/user")
	categoryRouter := rootRouter.Group("/category")
	productRouter := rootRouter.Group("/products")
	cartRouter := rootRouter.Group("/cart")
	orderRouter := rootRouter.Group("/orders")

	// Repositories
	roleRepo := role.NewRoleRepository(DB)
	roleRepo.Migration()
	roleRepo.InserSampleData()
	
	userRepo := users.NewUserRepository(DB)
	userRepo.Migration()
	categoryRepo := category.NewCategoryRepository(DB)
	categoryRepo.Migration()
	productRepo := product.NewProductRepository(DB)
	productRepo.Migration()
	cartRepo := cart.NewCartRepository(DB)
	cartRepo.Migration()
	cartItemRepo := cartItem.NewCartCartItemRepository(DB)
	cartItemRepo.Migration()
	orderItemRepo := orderItem.NewOrderItemRepository(DB)
	orderItemRepo.Migration()
	orderRepo := order.NewOrderRepository(DB)
	orderRepo.Migration()

	// Services
	categoryService := category.NewCategoryService(categoryRepo)
	productService := product.NewProductService(productRepo)
	cartService := cart.NewCartService(cartRepo, productRepo, cartItemRepo)
	userService := users.NewUserService(userRepo, cartService)
	orderService := order.NewOrderService(orderRepo, orderItemRepo, cartService, productService, cartRepo)

	// Handlers
	users.NewAuthHandler(authRouter, cfg, userService)
	category.NewCategoryHandler(categoryRouter, cfg, categoryService)
	product.NewProductHandler(productRouter, cfg, productService)
	cart.NewCartHandler(cartRouter, cfg, cartService)
	order.NewOrderHandler(orderRouter, cfg, orderService, cartService)


	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	r.GET("healthx", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	r.GET("readyx", func(c *gin.Context) {
		db, err := DB.DB()
		if err != nil {
			zap.L().Fatal("Cannot get sql database instance ", zap.Error(err))
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		if err := db.Ping(); err != nil {
			zap.L().Fatal("Cannot ping database ", zap.Error(err))
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		c.JSON(http.StatusOK, nil)
	})

	log.Println("Shopping Cart service started!")

	graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int(time.Second)))
}
