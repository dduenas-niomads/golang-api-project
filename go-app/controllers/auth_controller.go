package controllers

import (
	"go-app/initializers"
	"go-app/data/requests"
	// "go-app/data/responses"
	"go-app/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	"os"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	
}

func Register(c *gin.Context) {

	db := initializers.DatabaseConnection()
	userRegisterRequest := requests.UserRegistrationRequest{}

	if err := c.ShouldBindJSON(&userRegisterRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User

	db.Where("email=?", userRegisterRequest.Email).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists."})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name: userRegisterRequest.Name,
		Email: userRegisterRequest.Email,
		Password: string(passwordHash),
	}

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func Login(c *gin.Context) {

	db := initializers.DatabaseConnection()
	userLoginRequest := requests.UserLoginRequest{}

	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var userFound models.User
	db.Where("email=?", userLoginRequest.Email).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found."})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userLoginRequest.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password."})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate token."})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func GetUserProfile(c *gin.Context) {

	user, _ := c.Get("currentUser")

	c.JSON(200, gin.H{
		"user": user,
	})
}