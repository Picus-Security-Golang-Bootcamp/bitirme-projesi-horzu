package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpErr "github.com/horzu/golang/cart-api/internal/httpErrors"
)

type orderHandler struct {
	repo *OrderRepository
}

func NewOrderHandler(r *gin.RouterGroup, repo *OrderRepository) {
	h := &orderHandler{repo: repo}

	r.GET("/", h.getAll)
}

func (o *orderHandler) getAll(c *gin.Context) {
	orders, err := o.repo.getAll()
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, OrdersToResponse(orders))
}


