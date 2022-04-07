package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/horzu/golang/cart-api/internal/api"
	httpErr "github.com/horzu/golang/cart-api/internal/httpErrors"
	"github.com/horzu/golang/cart-api/pkg/config"
	jwtHelper "github.com/horzu/golang/cart-api/pkg/jwt"
)

type authHandler struct {
	cfg *config.Config
}

func NewAuthHandler(r *gin.RouterGroup, cfg *config.Config){
	a := authHandler{cfg:cfg}

	r.POST("/login", a.login)
}

func (a *authHandler) login(c *gin.Context){
	var req api.Login
	if err:=c.Bind(&req); err!=nil{
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "Check your request body", nil)))
		return
	}

	user := GetUser(*req.Email, *req.Password)
	if user == nil{
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "user not found", nil)))
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Id,
		"email": user.Email,
		"iat": time.Now().Unix(),
		"iss": os.Getenv("ENV"),
		"exp": time.Now().Add(24* time.Hour).Unix(),
		"role": user.Roles,
	})

	token := jwtHelper.GenerateToken(jwtClaims, a.cfg.JWTConfig.SecretKey)

	c.JSON(http.StatusOK, token)
}
