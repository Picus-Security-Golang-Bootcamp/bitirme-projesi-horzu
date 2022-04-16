package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horzu/golang/cart-api/internal/httpErrors"
	"github.com/horzu/golang/cart-api/pkg/config"
	mw "github.com/horzu/golang/cart-api/pkg/middleware"
	"github.com/horzu/golang/cart-api/pkg/pagination"
)

type categoryHandler struct {
	cfg     *config.Config
	service Service
}

func NewCategoryHandler(r *gin.RouterGroup, cfg *config.Config, service Service) {
	h := &categoryHandler{service: service,
		cfg: cfg}
	r.Use(mw.AdminAuthMiddleware(cfg.JWTConfig.SecretKey))
	r.POST("/create", h.createBulk)
	r.GET("/", h.listAllCategories)

}

func (p *categoryHandler) createBulk(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "Error reading file",
		})
		return
	}
	
	_, err = p.service.CreateBulk(c.Request.Context(), file)

	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "records created!",
	})
}

func (p *categoryHandler) listAllCategories(c *gin.Context) {
	page, pageSize := pagination.GetPaginationParametersFromRequest(c)
	categories, count, err := p.service.ListAll(c.Request.Context(), page, pageSize)

	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	result := pagination.NewPaginatedResponse(c, int(count))
	result.Items = categories

	c.JSON(http.StatusOK, result)
}
