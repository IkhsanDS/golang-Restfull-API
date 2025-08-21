package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"github.com/IkhsanDS/golang-api/auth"
	"github.com/IkhsanDS/golang-api/database"
	"github.com/IkhsanDS/golang-api/models"
)

var validate = validator.New()

type RegisterInput struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// Register godoc
// @Summary      Register new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload body RegisterInput true "Register payload"
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} gin.H
// @Router       /auth/register [post]
func Register(c *gin.Context) {
	var in RegisterInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	u := models.User{Name: in.Name, Email: in.Email, Password: string(hash), Role: "user"}
	if err := database.DB.Create(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already used"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": u.ID, "email": u.Email, "name": u.Name})
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login godoc
// @Summary      Login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload body LoginInput true "Login payload"
// @Success      200 {object} TokenResponse
// @Failure      401 {object} gin.H
// @Router       /auth/login [post]
func Login(c *gin.Context) {
	var in LoginInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var u models.User
	if err := database.DB.Where("email = ?", in.Email).First(&u).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	at, _, _ := auth.GenerateToken(u.ID, u.Email, u.Role, u.TokenVersion, 15*time.Minute)
	rt, _, _ := auth.GenerateToken(u.ID, u.Email, u.Role, u.TokenVersion, 7*24*time.Hour)
	c.JSON(http.StatusOK, TokenResponse{AccessToken: at, RefreshToken: rt})
}

// Me godoc
// @Summary      Get current user
// @Tags         auth
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} models.User
// @Failure      401 {object} gin.H
// @Router       /auth/me [get]
func Me(c *gin.Context) {
	email, _ := c.Get("email")
	var u models.User
	if err := database.DB.Where("email = ?", email.(string)).First(&u).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not found"})
		return
	}
	u.Password = ""
	c.JSON(http.StatusOK, u)
}
