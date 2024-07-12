package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/services"
)

type UserHandler interface {
	GetUsers(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserHandlerImpl struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{service: service}
}

func (h *UserHandlerImpl) GetUsers(c *gin.Context) {
	// Call Users service
	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h *UserHandlerImpl) GetUser(c *gin.Context) {
	// Call User service
	user, err := h.service.GetUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *UserHandlerImpl) CreateUser(c *gin.Context) {
	// Parse body request to User model
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid request body: %s", err.Error()),
		})
		return
	}

	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return response
	c.JSON(http.StatusCreated, gin.H{
		"data": createdUser,
	})
}

func (h *UserHandlerImpl) UpdateUser(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	updatedUser, err := h.service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": updatedUser,
	})
}

func (h *UserHandlerImpl) DeleteUser(c *gin.Context) {
	deletedUser, err := h.service.DeleteUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": deletedUser,
	})
}
