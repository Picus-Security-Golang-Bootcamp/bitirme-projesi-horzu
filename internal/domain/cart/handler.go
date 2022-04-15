package cart

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/horzu/golang/cart-api/internal/api"
	httpErr "github.com/horzu/golang/cart-api/internal/httpErrors"
	"github.com/horzu/golang/cart-api/pkg/config"
	jwtHelper "github.com/horzu/golang/cart-api/pkg/jwt"
)

type cartHandler struct {
	service Service
	cfg     *config.Config
}

func NewCartHandler(r *gin.RouterGroup, cfg *config.Config, service Service) {
	h := &cartHandler{service: service, cfg: cfg}

	r.GET("/", h.listCartItems)
	r.POST("/:id", h.createCart)
	
	r.POST("/item", h.addTocart)
	r.PUT("/:id/item/:itemId/quantity/:quantity", h.updateItem)
	r.DELETE("/item/:itemId", h.deleteItem)

}

func (c *cartHandler) listCartItems(g *gin.Context) {
	id := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.cfg.JWTConfig.SecretKey)

	result, err := c.service.Get(g.Request.Context(), id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		g.JSON(http.StatusNotFound, "")
	}
	g.JSON(http.StatusOK, result)
}

func (c *cartHandler) createCart(g *gin.Context) {
	id := g.Param("id")

	if b, err := c.service.Create(g.Request.Context(), id); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	} else {

		g.JSON(http.StatusCreated, map[string]string{"id": b.Id})
	}
}

func (c *cartHandler) addTocart(g *gin.Context) {
	// "userid'yi auth'dan çek sonra da aktif cartı db'den çek"
	var req *api.CartItem
	if err := g.Bind(&req); err != nil {
		g.JSON(httpErr.ErrorResponse(httpErr.CannotBindGivenData))
		return
	}

	if err := req.Validate(strfmt.NewFormats()); err != nil {
		g.JSON(httpErr.ErrorResponse(err))
	}

	if itemId, err := c.service.AddItem(g.Request.Context(), *&req.Product.Sku, req.CartID, *&req.Quantity); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	} else {
		g.JSON(http.StatusCreated, map[string]string{"id": itemId})
	}
}

func (c *cartHandler) updateItem(g *gin.Context) {
	cartId := g.Param("id")
	itemId := g.Param("itemId")
	quantity, err := strconv.Atoi(g.Param("quantity"))

	if len(cartId) == 0 || len(itemId) == 0 || err != nil || quantity <= 0 {
		g.JSON(http.StatusBadRequest, "Failed to update item. BasketId or BasketItem Id is null or empty.")
	}
	if err := c.service.UpdateItem(g.Request.Context(), cartId, itemId, uint(quantity)); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	}
	g.JSON(http.StatusAccepted, "item updated")

}

func (c *cartHandler) deleteItem(g *gin.Context) {
	id := "userid'yi auth'dan çek"
	itemId := g.Param("itemId")

	if err := c.service.DeleteItem(g.Request.Context(), id, itemId); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	}
	g.JSON(http.StatusAccepted, "")
}


func getUserIdFromAuthToken(token, secretKey string) string {
	decodedClaims := jwtHelper.VerifyToken(token, secretKey)
	userId := decodedClaims.UserId

	return userId
}