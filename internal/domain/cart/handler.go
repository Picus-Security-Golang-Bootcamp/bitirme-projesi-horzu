package cart

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/httpErrors"
	"github.com/horzu/golang/cart-api/pkg/config"
	mw "github.com/horzu/golang/cart-api/pkg/middleware"
)

type cartHandler struct {
	service Service
	cfg     *config.Config
}

func NewCartHandler(r *gin.RouterGroup, cfg *config.Config, service Service) {
	h := &cartHandler{service: service, cfg: cfg}

	r.Use(mw.UserAuthMiddleware(cfg.JWTConfig.SecretKey))
	r.GET("/", h.listCartItems)

	r.POST("/item", h.addTocart)
	r.PUT("/item/:itemId/quantity/:quantity", h.updateItem)
	r.DELETE("/item/:itemId", h.deleteItem)

}

func (c *cartHandler) listCartItems(g *gin.Context) {
	userId := g.GetString("userID")

	result, err := c.service.GetAllCartItems(g, userId)
	if err != nil {
		g.JSON(http.StatusInternalServerError, httpErrors.ParseErrors(err))
	}
	if result == nil {
		g.JSON(http.StatusNotFound, "")
	}
	g.JSON(http.StatusOK, api.SuccessfulAPIResponse{
		Code: http.StatusOK,
		Message: "List of cart items",
		Data: result,
	})
}

func (c *cartHandler) addTocart(g *gin.Context) {
	userId := g.GetString("userID")

	sku := g.Query("itemId")

	quantity, err := strconv.Atoi(g.Query("quantity"))

	if err != nil {
		g.JSON(http.StatusBadRequest, httpErrors.ParseErrors(err))
		return
	}

	cart, err := c.service.GetCartByUserId(g, userId)

	if itemId, err := c.service.AddItem(g.Request.Context(), sku, cart.Id, int64(quantity)); err != nil {
		g.JSON(http.StatusBadRequest, httpErrors.ParseErrors(err))
		return
	} else {
		g.JSON(http.StatusCreated, api.SuccessfulAPIResponse{
				Code: http.StatusAccepted,
				Message: "Product added to cart",
				Data: map[string]string{"Added Product Id": itemId},
			})
		return

	}
}

func (c *cartHandler) updateItem(g *gin.Context) {
	userId := g.GetString("userID")

	itemId := g.Param("itemId")
	quantity, err := strconv.Atoi(g.Param("quantity"))

	cart, err := c.service.GetCartByUserId(g, userId)

	if len(cart.Id) == 0 || len(itemId) == 0 || err != nil || quantity <= 0 {
		g.JSON(http.StatusBadRequest, "Failed to update item. CartId or CartItem Id is null or empty.")
		return
	}
	if err := c.service.UpdateItem(g.Request.Context(), cart.Id, itemId, uint(quantity)); err != nil {
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}
	g.JSON(http.StatusAccepted, api.SuccessfulAPIResponse{
		Code: http.StatusAccepted,
		Message: "item updated",
	})

}

func (c *cartHandler) deleteItem(g *gin.Context) {
	userId := g.GetString("userID")

	itemId := g.Param("itemId")

	cart, err := c.service.GetCartByUserId(g, userId)

	if err != nil {
		g.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.service.DeleteItem(g.Request.Context(), cart.Id, itemId); err != nil {
		g.JSON(http.StatusBadRequest, httpErrors.ParseErrors(err))
		return
	}
	g.JSON(http.StatusAccepted, api.SuccessfulAPIResponse{
		Code: http.StatusAccepted,
		Message: "item deleted from the cart",
	})
}
