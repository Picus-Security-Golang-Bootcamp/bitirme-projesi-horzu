package orderProduct

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/httpErrors"
)

type orderProductHandler struct {
	repo *ProductRepository
}

func NewOrderProductHandler(r *gin.RouterGroup, repo *OrderProductRepository){
	h := &orderProductHandler{repo: repo}

	r.POST("/create", h.create)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

func (p *orderProductHandler) create(c *gin.Context){
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

func (p *orderProductHandler) getByID(c *gin.Context){
	product, err := p.repo.getByID(c.Param("id"))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, productToResponse(product))
}

func (p *orderProductHandler) update(c *gin.Context){
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

	product, err := p.repo.update(responseToProduct(productBody))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
	}

	c.JSON(http.StatusOK, productToResponse(product))
}

func (p *orderProductHandler) delete(c *gin.Context){
	err := p.repo.delete(c.Param("id"))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}