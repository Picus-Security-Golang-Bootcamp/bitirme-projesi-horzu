package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/horzu/golang/cart-api/internal/api"
	httpErr "github.com/horzu/golang/cart-api/internal/httpErrors"
)

type cartHandler struct {
	repo *CartRepository
}

func NewCartHandler(r *gin.RouterGroup, repo *CartRepository) {
	h := &cartHandler{repo: repo}

	r.GET("/", h.getAll)
	r.POST("/create", h.create)
	r.GET("/:id", h.getByID)
}

func (o *cartHandler) create(c *gin.Context){
	cart := api.Cart{}
	if err := c.Bind(&cart); err !=nil{
		c.JSON(httpErr.ErrorResponse(httpErr.CannotBindGivenData))
		return
	}

	if err := cart.Validate(strfmt.NewFormats()); err!=nil{
		c.JSON(httpErr.ErrorResponse(err))
	}

	// newcart, err := o.repo.create(responseToCart(&cart))
	// if err!=nil{
	// 	c.JSON(httpErr.ErrorResponse(err))
	// }

	// c.JSON(http.StatusOK, cartToResponse(newcart))
}

func (o *cartHandler) getAll(c *gin.Context) {
	// cart, err := o.repo.getAll()
	// if err != nil {
	// 	c.JSON(httpErr.ErrorResponse(err))
	// 	return
	// }

	// c.JSON(http.StatusOK, cartsToResponse(cart))
}

func (o *cartHandler) getByID(c *gin.Context){
	// cart, err := o.repo.getByID(c.Param("id"))
	// if err!=nil{
	// 	c.JSON(httpErr.ErrorResponse(err))
	// 	return
	// }
	// c.JSON(http.StatusOK, cartToResponse(cart))
}