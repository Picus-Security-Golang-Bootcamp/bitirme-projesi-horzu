package category

import (
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/httpErrors"
)

type categoryHandler struct {
	repo *CategoryRepository
}

func NewCategoryHandler(r *gin.RouterGroup, repo *CategoryRepository) {
	h := &categoryHandler{repo: repo}

	r.POST("/createbulkcategory", h.createBulk)
	r.GET("/listcategory", h.listAllCategories)

}

func (p *categoryHandler) createBulk(c *gin.Context) {
	categoryBody := &api.Category{}

	file, err := c.FormFile("file")
	filedata, _ := file.Open()
	defer filedata.Close()


	csvLines, err := csv.NewReader(filedata).ReadAll()
	
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  err,
		})
	}

	// Serializer yapılacak!! responseToCategory
	for _, record := range csvLines[1:] {
		categoryBody.Name = &record[0]
		categoryBody.Image = &record[1]
		p.repo.create(responseToCategory(categoryBody))

	}
}

func (p *categoryHandler) listAllCategories(c *gin.Context) {
	categories, err := p.repo.getAll()
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, categoriesToResponse(categories))
}

