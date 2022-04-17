package order

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/domain/cart"
	"github.com/horzu/golang/cart-api/pkg/config"
	mw "github.com/horzu/golang/cart-api/pkg/middleware"
)

type orderHandler struct {
	service Service
	cfg     *config.Config
	cartService *cart.CartService
}

func NewOrderHandler(r *gin.RouterGroup, cfg *config.Config, service Service, cartService *cart.CartService) {
	h := &orderHandler{service: service, cfg: cfg, cartService: cartService}
	r.Use(mw.UserAuthMiddleware(cfg.JWTConfig.SecretKey))

	r.POST("/", h.CompleteOrderWithUserId)
	r.GET("/", h.ListAll)

}

func (order *orderHandler) CompleteOrderWithUserId(g *gin.Context) {
	userId := g.GetString("userID")
	fmt.Println(userId)
	err := order.service.CompleteOrderWithUserId(g, userId)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, api.ErrorAPIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusCreated, api.SuccessfulAPIResponse{
		Code:    http.StatusOK,
		Message: "ok",
	})
}

func (order *orderHandler) ListAll(g *gin.Context) {
	userId := g.GetString("userID")

	orders, err:= order.service.GetAll(g, userId)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, api.ErrorAPIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusCreated, api.SuccessfulAPIResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Data: orders,
	})
}

