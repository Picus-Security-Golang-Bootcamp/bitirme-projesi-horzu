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
	if err!=nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "Error!!!",
		})
		return 
	}
	p.service.CreateBulk(c.Request.Context(), file)

	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"status": "Cretead!!",
	})
}

func (p *categoryHandler) listAllCategories(c *gin.Context) {
	page := pagination.NewFromGinRequest(c, -1)
	categories, err := p.service.ListAll(c.Request.Context(), page)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, categoriesToResponse(&categories))
}
