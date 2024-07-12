package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/handlers"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/middlewares"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/pkg/database/mongodb"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/repositories"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/services"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/token"
)

func SetupRoutes(router *gin.Engine) {
	// Initialize the user handler
	userRepo := repositories.NewUserRepository(mongodb.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Initialize the job handler
	jobRepo := repositories.NewJobRepository(mongodb.DB)
	jobService := services.NewJobService(jobRepo)
	jobHandler := handlers.NewJobHandler(jobService)

	// Initialize the survey handler
	surveyRepo := repositories.NewSurveyRepository(mongodb.DB)
	surveyService := services.NewSurveyService(surveyRepo)
	surveyHandler := handlers.NewSurveyHandler(surveyService)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Berufsvernetzen!",
		})
	})

	// Add /api/v1 to the base path
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello Berufsvernetzen!",
			})
		})

		v1.POST("/seed", middlewares.AuthMiddleware(token.TokenMaker), handlers.SeedingDatebase)

		// AuthRoute
		auth := v1.Group("/auth")
		{
			auth.POST("/login", handlers.Login)
			auth.POST("/register", handlers.Register)
			auth.POST("/verify", handlers.VerifyHandler)
		}

		// UsersRoute
		users := v1.Group("/users")
		{
			users.GET("/", userHandler.GetUsers)
			users.GET("/:id", middlewares.AuthMiddleware(token.TokenMaker), userHandler.GetUser)
			users.POST("/", userHandler.CreateUser)
			users.PATCH("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// JobsRoute
		jobs := v1.Group("/jobs")
		{
			jobs.GET("/", jobHandler.GetJobs)
			jobs.GET("/:id", jobHandler.GetJob)
			jobs.POST("/search", jobHandler.SearchJobs)
		}

		surveys := v1.Group("/surveys")
		{
			surveys.GET("/", surveyHandler.GetSurveys)
			surveys.GET("/:id", surveyHandler.GetSurvey)
			surveys.POST("/", surveyHandler.CreateSurvey)
			surveys.PATCH("/:id", surveyHandler.UpdateSurvey)
			surveys.DELETE("/:id", surveyHandler.DeleteSurvey)
			surveys.POST("/json", surveyHandler.PostSurveyJson)
			surveys.POST("/jsons", surveyHandler.PostSurveyJsons)
			surveys.GET("/feedbacks", surveyHandler.GetFeedbacks)
		}

		// AnalyticsRoute
		analytics := v1.Group("/analytics")
		{
			analytics.POST("/generate", handlers.GenerateAnalytics)
		}
	}
}
