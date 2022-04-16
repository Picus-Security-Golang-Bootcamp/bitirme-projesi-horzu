package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/httpErrors"
	"github.com/horzu/golang/cart-api/pkg/config"
	mw "github.com/horzu/golang/cart-api/pkg/middleware"
	"github.com/horzu/golang/cart-api/pkg/pagination"
)

type productHandler struct {
	service Service
	cfg     *config.Config

}

func NewProductHandler(r *gin.RouterGroup, cfg *config.Config, service Service) {
	h := &productHandler{service: service, cfg: cfg}


	r.POST("/create", h.create).Use(mw.AdminAuthMiddleware(cfg.JWTConfig.SecretKey))
	r.PUT("/:sku", h.update).Use(mw.AdminAuthMiddleware(cfg.JWTConfig.SecretKey))
	r.DELETE("/:sku", h.delete).Use(mw.AdminAuthMiddleware(cfg.JWTConfig.SecretKey))
	r.GET("/", h.listProduct)
	r.GET("/search/:sku", h.searchProduct)
}

func (p *productHandler) create(c *gin.Context) {
	var productBody *api.ProductCreateProductRequest

	if err := c.Bind(productBody); err != nil {
		c.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
		return
	}

	if err := productBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	productCreated := responseToProduct(productBody)

	err := p.service.CreateProduct(c.Request.Context(), productCreated.Name, productCreated.Description, int64(productBody.Stock), *&productBody.Price, productCreated.CategoryId)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, &productCreated)
}

func (p *productHandler) update(c *gin.Context) {
	sku := c.Param("sku")
	productBody := &api.ProductUpdateProductRequest{Sku: sku}
	if err := c.Bind(&productBody); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	if err := productBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	err := p.service.UpdateProduct(c.Request.Context(), responseToUpdateProduct(productBody))
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	c.JSON(http.StatusOK, "Product Updated")
}

func (p *productHandler) delete(c *gin.Context) {
	err := p.service.DeleteProduct(c.Request.Context(), c.Param("sku"))
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusNoContent, "Product deleted")
}

func (p *productHandler) listProduct(c *gin.Context) {
	page := pagination.NewFromGinRequest(c, -1)
	result := p.service.GetAll(c.Request.Context(), page)
	if result == nil {
		c.JSON(http.StatusInternalServerError,"No products found")
		return
	}


	c.JSON(http.StatusOK, page)
}

func (p *productHandler) searchProduct(c *gin.Context) {
	page := pagination.NewFromGinRequest(c, -1)
	//get search keyword from query
	searchItem := c.Param("sku")
	result := p.service.SearchProduct(c.Request.Context(), searchItem, page)
	if result == nil {
		c.JSON(http.StatusInternalServerError,"No products found")
		return
	}

	c.JSON(http.StatusOK, page)
}
