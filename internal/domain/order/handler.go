package order

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/domain/cart"
	"github.com/horzu/golang/cart-api/internal/httpErrors"
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

	r.POST("/", h.completeOrderWithUserId)
	r.GET("/", h.listAll)
	r.DELETE("/:id", h.cancelOrder)

}

func (order *orderHandler) completeOrderWithUserId(g *gin.Context) {
	userId := g.GetString("userID")
	fmt.Println(userId)
	err := order.service.CompleteOrderWithUserId(g, userId)

	if err != nil {
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}

	g.JSON(http.StatusCreated, api.SuccessfulAPIResponse{
		Code:    http.StatusOK,
		Message: "Order Created",
	})
}

func (order *orderHandler) listAll(g *gin.Context) {
	userId := g.GetString("userID")

	orders, err:= order.service.GetAll(g, userId)

	if err != nil {
		log.Println(err.Error())
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}

	g.JSON(http.StatusCreated, api.SuccessfulAPIResponse{
		Code:    http.StatusOK,
		Message: "All orders listed",
		Data: orders,
	})
}

func (order *orderHandler) cancelOrder(g *gin.Context) {
	userId := g.GetString("userID")

	orderId := g.Param("id")

	err:= order.service.CancelOrder(g, userId, orderId)

	if err != nil {
		log.Println(err.Error())
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}

	g.JSON(http.StatusCreated, api.SuccessfulAPIResponse{
		Code:    http.StatusOK,
		Message: "Order Canceled",
	})
}
