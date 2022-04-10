package productCategory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/httpErrors"
)

type productCategoryHandler struct {
	repo *ProductCategoryRepository
}

func NewProductCategoryHandler(r *gin.RouterGroup, repo *ProductCategoryRepository){
	h := &productCategoryHandler{repo: repo}

	r.POST("/create", h.create)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete)
}

func (p *productCategoryHandler) create(c *gin.Context){
	categoryBody := &api.Category{}
		
	if err:= c.Bind(categoryBody); err!=nil{
		c.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
		return
	}

	if err:= categoryBody.Validate(strfmt.NewFormats()); err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	product, err := p.repo.create(responseToCategory(categoryBody))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, &product)
}

func (p *productCategoryHandler) getByID(c *gin.Context){
	category, err := p.repo.getByID(c.Param("id"))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, categoryToResponse(category))
}

func (p *productCategoryHandler) update(c *gin.Context){
	id := c.Param("id")
	categoryBody := &api.Category{ID: id}
	if err:=c.Bind(&categoryBody); err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	if err:= categoryBody.Validate(strfmt.NewFormats()); err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	product, err := p.repo.update(responseToCategory(categoryBody))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
	}

	c.JSON(http.StatusOK, categoryToResponse(product))
}

func (p *productCategoryHandler) delete(c *gin.Context){
	err := p.repo.delete(c.Param("id"))
	if err!=nil{
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}