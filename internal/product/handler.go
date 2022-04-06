package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/httpErrors"
)

type productHandler struct {
	repo *ProductRepository
}

func NewProductHandler(r *gin.RouterGroup, repo *ProductRepository){
	h := &productHandler{repo: repo}

	r.POST("/create", h.create)
	r.GET("/:id", h.getByID)
}

func (p *productHandler) create(c *gin.Context){
	productBody := &api.Product{}
	if err:= c.Bind(productBody); err!=nil{
		c.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
		return
	}

	if err:= productBody.Validate(strfmt.NewFormats()); err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	product, err := p.repo.create(responseToProduct(productBody))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, &product)
}

func (p *productHandler) getByID(c *gin.Context){
	product, err := p.repo.getByID(c.Param("id"))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, productToResponse(product))
}