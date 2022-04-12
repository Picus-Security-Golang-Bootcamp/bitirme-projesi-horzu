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

	r.POST("/create", h.Create)
	r.GET("/:id", h.GetByID)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}

func (p *productHandler) Create(c *gin.Context){
	productBody := &api.Product{}
		
	if err:= c.Bind(productBody); err!=nil{
		c.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
		return
	}

	if err:= productBody.Validate(strfmt.NewFormats()); err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	product, err := p.repo.Create(responseToProduct(productBody))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, &product)
}

func (p *productHandler) GetByID(c *gin.Context){
	product, err := p.repo.GetByID(c.Param("id"))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, productToResponse(product))
}

func (p *productHandler) Update(c *gin.Context){
	id := c.Param("id")
	productBody := &api.Product{ID: id}
	if err:=c.Bind(&productBody); err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	if err:= productBody.Validate(strfmt.NewFormats()); err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	product, err := p.repo.Update(responseToProduct(productBody))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
	}

	c.JSON(http.StatusOK, productToResponse(product))
}

func (p *productHandler) Delete(c *gin.Context){
	err := p.repo.Delete(c.Param("id"))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}