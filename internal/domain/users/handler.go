package users

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/httpErrors"
	httpErr "github.com/horzu/golang/cart-api/internal/httpErrors"
	"github.com/horzu/golang/cart-api/pkg/config"
	jwtHelper "github.com/horzu/golang/cart-api/pkg/jwt"
	mw "github.com/horzu/golang/cart-api/pkg/middleware"
)

type authHandler struct {
	cfg  *config.Config
	repo *UserRepository
}

func NewAuthHandler(r *gin.RouterGroup, cfg *config.Config, repo *UserRepository) {
	a := authHandler{cfg: cfg,
		repo: repo}

	r.POST("/login", a.login)
	r.POST("/signup", a.Signup)

	r.Use(mw.AdminAuthMiddleware(cfg.JWTConfig.SecretKey))
	r.POST("/decode", a.VerifyToken)
}

type tokenStruct struct {
	token        string
	refreshToken string
}

func (a *authHandler) Signup(c *gin.Context) {
	user := User{}

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := a.repo.SaveUser(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email is already registered."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success!"})
}

func (a *authHandler) login(c *gin.Context) {
	var req api.UserCreateUserRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "Check your request body", nil)))
		return
	}

	user, err := a.repo.LoginCheck(*req.Email, *req.Password)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}
	if user == nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "user not found", nil)))
	}

	jwtClaimsForToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"role":  user.Role.Role,
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Duration(a.cfg.JWTConfig.SessionTime) * time.Second).Unix(),
	})
	jwtClaimsForRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role":  user.Role.Role,
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Duration(a.cfg.JWTConfig.SessionTime*168) * time.Second).Unix(),
	})

	token := jwtHelper.GenerateToken(jwtClaimsForToken, a.cfg.JWTConfig.SecretKey)
	refreshToken := jwtHelper.GenerateToken(jwtClaimsForRefreshToken, a.cfg.JWTConfig.SecretKey)

	tokens := &tokenStruct{
		token:        token,
		refreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, &tokens.token)
}

func (a *authHandler) VerifyToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	decodeClaims := jwtHelper.VerifyToken(token, a.cfg.JWTConfig.SecretKey)

	c.JSON(http.StatusOK, decodeClaims)
}
