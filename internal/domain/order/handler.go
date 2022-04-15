package order

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horzu/golang/cart-api/pkg/config"
	jwtHelper "github.com/horzu/golang/cart-api/pkg/jwt"
	mw "github.com/horzu/golang/cart-api/pkg/middleware"
)

type orderHandler struct {
	service Service
	cfg     *config.Config
}

func NewOrderHandler(r *gin.RouterGroup, cfg *config.Config, service Service) {
	h := &orderHandler{service: service, cfg: cfg}
	r.Use(mw.AdminAuthMiddleware(cfg.JWTConfig.SecretKey))

	r.POST("/", h.CompleteOrderWithUserId)

}

func (c *orderHandler) CompleteOrderWithUserId(g *gin.Context) {
	userId := getUserIdFromAuthToken(g.GetHeader("Authorization"), c.cfg.JWTConfig.SecretKey)

	err := c.service.CompleteOrderWithUserId(g.Request.Context(), userId)

	if err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, ApiErrorResponse{
			IsSuccess:    false,
			ErrorMessage: err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusCreated, ApiOkResponse{
		IsSuccess: true,
		Message:   "ok",
	})
}

func getUserIdFromAuthToken(token, secretKey string) string {
	decodedClaims := jwtHelper.VerifyToken(token, secretKey)
	userId := decodedClaims.UserId

	return userId
}

type ApiErrorResponse struct {
	IsSuccess    bool   `json:"is_success"`
	ErrorMessage string `json:"error_message"`
}

type ApiOkResponse struct {
	IsSuccess bool        `json:"is_success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
