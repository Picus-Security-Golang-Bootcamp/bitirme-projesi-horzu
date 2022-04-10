package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/httpErrors"
	httpErr "github.com/horzu/golang/cart-api/internal/httpErrors"
	"github.com/horzu/golang/cart-api/internal/models"
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

	r.Use(mw.AuthMiddleware(cfg.JWTConfig.SecretKey))
	r.POST("/decode", a.VerifyToken)
}


func (a *authHandler) Signup(c *gin.Context){
	var input *api.User

	if err:= c.ShouldBindJSON(&input);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	user.Email = input.Email
	user.Password = input.Password
	user.IsAdmin = input.IsAdmin

	_, err := a.repo.SaveUser(&user)

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "This email is already registered."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success!"})
}

func (a *authHandler) login(c *gin.Context) {
	var req api.Login
	if err := c.Bind(&req); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "Check your request body", nil)))
		return
	}

	u := models.User{}
	u.Email = *req.Email
	u.Password = *req.Password

	user, err := a.repo.GetUser(u.Email, u.Password)
	fmt.Println(user)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}
	if user == nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "user not found", nil)))
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"isAdmin": user.IsAdmin,
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Duration(a.cfg.JWTConfig.SessionTime) * time.Second).Unix(),
	})

	token := jwtHelper.GenerateToken(jwtClaims, a.cfg.JWTConfig.SecretKey)

	c.JSON(http.StatusOK, token)
}

func (a *authHandler) VerifyToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	decodeClaims := jwtHelper.VerifyToken(token, a.cfg.JWTConfig.SecretKey)

	c.JSON(http.StatusOK, decodeClaims)
}
