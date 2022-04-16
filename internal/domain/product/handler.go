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


	r.GET("/", h.listProduct)
	r.GET("/search/:sku", h.searchProduct)
	r.Use(mw.AdminAuthMiddleware(cfg.JWTConfig.SecretKey))
	r.POST("/", h.create)
	r.PUT("/", h.update)
	r.DELETE("/:sku", h.delete)
}

func (p *productHandler) create(c *gin.Context) {
	var productBody *api.ProductCreateProductRequest

	if err := c.Bind(&productBody); err != nil {
		c.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
		return
	}

	if err := productBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	productCreated := requestToProduct(productBody)

	err := p.service.CreateProduct(c, productCreated.Name, productCreated.Description, int64(*productBody.Stock), *productBody.Price, productCreated.CategoryId)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, &productCreated)
}

func (p *productHandler) update(c *gin.Context) {
	var productBody *api.ProductUpdateProductRequest
	if err := c.Bind(&productBody); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	if err := productBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	err := p.service.UpdateProduct(c.Request.Context(), requestToUpdateProduct(productBody))
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

	c.JSON(http.StatusAccepted, api.SuccessfulAPIResponse{
		Code: http.StatusAccepted,
		Message: "Product Deleted",
	})
}

func (p *productHandler) listProduct(c *gin.Context) {
	page, pageSize := pagination.GetPaginationParametersFromRequest(c)
	products, count, err := p.service.GetAll(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	result := pagination.NewPaginatedResponse(c, int(count))
	result.Items = products

	c.JSON(http.StatusOK, result)
}

func (p *productHandler) searchProduct(c *gin.Context) {
	page, pageSize := pagination.GetPaginationParametersFromRequest(c)

	//get search keyword from query
	searchItem := c.Param("sku")
	products, count, err := p.service.SearchProduct(c.Request.Context(), searchItem, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError,"No products found")
		return
	}

	result := pagination.NewPaginatedResponse(c, int(count))
	result.Items = products

	c.JSON(http.StatusOK, result)
}
