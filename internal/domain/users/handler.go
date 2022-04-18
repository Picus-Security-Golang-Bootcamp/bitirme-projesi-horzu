package users

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/httpErrors"
	httpErr "github.com/horzu/golang/cart-api/internal/httpErrors"
	"github.com/horzu/golang/cart-api/pkg/config"
	jwtHelper "github.com/horzu/golang/cart-api/pkg/jwt"
	mw "github.com/horzu/golang/cart-api/pkg/middleware"
)

type authHandler struct {
	service Service
	cfg     *config.Config
}

type tokenStruct struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

var refreshTokenTime = 24 * 7

func NewAuthHandler(r *gin.RouterGroup, cfg *config.Config, service Service) {
	a := authHandler{service: service, cfg: cfg}

	r.POST("/login", a.login)
	r.POST("/signup", a.Signup)

	r.Use(mw.AdminAuthMiddleware(cfg.JWTConfig.SecretKey))
	r.POST("/decode", a.VerifyToken)
}

func (a *authHandler) Signup(c *gin.Context) {
	var user *api.UserCreateUserRequest

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if strings.Contains(*user.Password, " ")  {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "Password should not include empty strings", nil)))
		fmt.Println()
		return
	}

	if len(*user.Password) < 8 {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "Password should be minimum 8 characters", nil)))
		fmt.Println()
		return
	}


	if err := user.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	newUser, err := a.service.Create(c, responseToUser(user))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This user is already registered."})
		return
	}

	jwtClaimsForToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": newUser.Id,
		"role":   newUser.Role.Role,
		"email":  newUser.Email,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Duration(a.cfg.JWTConfig.SessionTime) * time.Second).Unix(),
	})

	jwtClaimsForRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": newUser.Id,
		"role":   newUser.Role.Role,
		"email":  newUser.Email,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Duration(a.cfg.JWTConfig.SessionTime*refreshTokenTime) * time.Second).Unix(),
	})

	token := jwtHelper.GenerateToken(jwtClaimsForToken, a.cfg.JWTConfig.SecretKey)
	refreshToken := jwtHelper.GenerateToken(jwtClaimsForRefreshToken, a.cfg.JWTConfig.SecretKey)

	c.JSON(http.StatusOK, tokenStruct{
		Token:        token,
		RefreshToken: refreshToken,
	})
}

func (a *authHandler) login(c *gin.Context) {
	var req api.Login
	if err := c.Bind(&req); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	if err := req.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	user, err := a.service.LoginCheck(*req.Email, *req.Password)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}
	if user == nil {
		c.JSON(httpErr.ErrorResponse(err))
	}

	jwtClaimsForToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.Id,
		"role":   user.Role.Role,
		"email":  user.Email,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Duration(a.cfg.JWTConfig.SessionTime) * time.Second).Unix(),
	})

	jwtClaimsForRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.Id,
		"role":   user.Role.Role,
		"email":  user.Email,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Duration(a.cfg.JWTConfig.SessionTime*refreshTokenTime) * time.Second).Unix(),
	})

	token := jwtHelper.GenerateToken(jwtClaimsForToken, a.cfg.JWTConfig.SecretKey)
	refreshToken := jwtHelper.GenerateToken(jwtClaimsForRefreshToken, a.cfg.JWTConfig.SecretKey)

	c.JSON(http.StatusOK, tokenStruct{
		Token:        token,
		RefreshToken: refreshToken,
	})
}

func (a *authHandler) VerifyToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	decodeClaims := jwtHelper.VerifyToken(token, a.cfg.JWTConfig.SecretKey)

	c.JSON(http.StatusOK, decodeClaims)
}
