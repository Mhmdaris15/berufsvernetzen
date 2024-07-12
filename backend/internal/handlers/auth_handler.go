package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/configs"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/pkg/database/mongodb"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/token"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserRegisterRequest struct {
	Name     string `bson:"name" json:"name" binding:"required,min=4,max=32"`
	Username string `bson:"username" json:"username" binding:"required,min=4,max=16"`
	Email    string `bson:"email" json:"email" binding:"required,email"`
	Password string `bson:"password" json:"password" binding:"required,min=8"`
}

type UserAuthResponse struct {
	Name     string `bson:"name" json:"name"`
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Token    string `bson:"token" json:"token"`
}

func Register(c *gin.Context) {
	var user UserRegisterRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Generate hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 16)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	UserCollection := mongodb.GetCollection(mongodb.DB, "Users")

	// Check if user with the same username or email already exists
	var result bson.M
	err = UserCollection.FindOne(c.Request.Context(), bson.M{"$or": []bson.M{
		{"username": user.Username},
		{"email": user.Email},
	}}).Decode(&result)

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
		return
	}

	var userReal models.User
	userReal.Name = user.Name
	userReal.Username = user.Username
	userReal.Email = user.Email
	userReal.Password = string(hashedPassword)

	// Insert user to database
	_, err = UserCollection.InsertOne(c.Request.Context(), userReal)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	generatedToken, err := token.TokenMaker.CreateToken(user.Username, configs.EnvAccessTokenDuration())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, UserAuthResponse{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Token:    generatedToken,
	})
}

type UserLoginEmailRequest struct {
	Email    string `bson:"email" json:"email" binding:"required,email"`
	Password string `bson:"password" json:"password" binding:"required,min=8"`
}

type UserLoginUsernameRequest struct {
	Username string `bson:"username" json:"username" binding:"required,min=4,max=16"`
	Password string `bson:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var user UserLoginEmailRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	UserCollection := mongodb.GetCollection(mongodb.DB, "Users")

	var result bson.M
	err := UserCollection.FindOne(c.Request.Context(), bson.M{"email": user.Email}).Decode(&result)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result["password"].(string)), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password is incorrect",
		})
		return
	}

	generatedToken, err := token.TokenMaker.CreateToken(result["username"].(string), configs.EnvAccessTokenDuration())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UserAuthResponse{
		Name:     result["name"].(string),
		Username: result["username"].(string),
		Email:    result["email"].(string),
		Token:    generatedToken,
	})
}

func VerifyHandler(c *gin.Context) {
	// Check if Authorization header exists
	if c.GetHeader("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header required",
		})
		return
	}

	// Get Token from request
	tokenString := c.GetHeader("Authorization")[len("Bearer "):]
	payload, err := token.TokenMaker.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("payload", payload)

	c.JSON(http.StatusOK, gin.H{
		"message": "Verified",
		"payload": payload,
	})
}
