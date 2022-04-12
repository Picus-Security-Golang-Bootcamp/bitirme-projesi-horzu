package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/horzu/golang/cart-api/internal/api"
	httpErr "github.com/horzu/golang/cart-api/internal/httpErrors"
)

type orderHandler struct {
	repo *OrderRepository
}

func NewOrderHandler(r *gin.RouterGroup, repo *OrderRepository) {
	h := &orderHandler{repo: repo}

	r.GET("/", h.getAll)
	r.POST("/create", h.create)
	r.GET("/:id", h.getByID)
}

func (o *orderHandler) create(c *gin.Context){
	orderBody := api.Order{}
	if err := c.Bind(&orderBody); err !=nil{
		c.JSON(httpErr.ErrorResponse(httpErr.CannotBindGivenData))
		return
	}

	if err := orderBody.Validate(strfmt.NewFormats()); err!=nil{
		c.JSON(httpErr.ErrorResponse(err))
	}

	order, err := o.repo.create(responseToOrder(&orderBody))
	if err!=nil{
		c.JSON(httpErr.ErrorResponse(err))
	}

	c.JSON(http.StatusOK, orderToResponse(*order))
}

func (o *orderHandler) getAll(c *gin.Context) {
	orders, err := o.repo.getAll()
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, ordersToResponse(orders))
}

func (o *orderHandler) getByID(c *gin.Context){
	order, err := o.repo.getByID(c.Param("id"))
	if err!=nil{
		c.JSON(httpErr.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, orderToResponse(*order))
}