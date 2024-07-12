package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/seeder"
)

func SeedingDatebase(c *gin.Context) {
	// handle seeding database logic

	users, surveys, err := seeder.Seed()
	if err != nil {
		if users == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else if surveys == nil {
			fmt.Errorf("Error when seeding database: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"users":   users,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Database seeded successfully",
		"users":   users,
		"surveys": surveys,
	})

}
