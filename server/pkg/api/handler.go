package api

import (
	"github.com/danielsobrado/ainovelprompter/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Handler struct {
	DB *gorm.DB
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	handler := &Handler{DB: db}

	router := gin.Default()

	// Configure CORS middleware
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3001"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(corsConfig))

	v1 := router.Group("/v1")
	{
		// Public routes
		v1.POST("/users/login", handler.Login)
		v1.POST("/users/register", handler.Register)

		if viper.GetBool("auth.enabled") {
			// Protected routes with authentication
			protected := v1.Group("")
			protected.Use(middleware.AuthMiddleware())
			{
				protected.GET("/users", handler.ListUsers)
				protected.POST("/users", handler.CreateUser)
				protected.GET("/users/:userId", handler.GetUser)
				protected.PUT("/users/:userId", handler.UpdateUser)
				protected.DELETE("/users/:userId", handler.DeleteUser)

				protected.GET("/trait-types", handler.ListTraitTypes)

				protected.GET("/texts", handler.ListTexts)
				protected.POST("/texts", handler.CreateText)
				protected.GET("/texts/:textId", handler.GetText)
				protected.PUT("/texts/:textId", handler.UpdateText)
				protected.DELETE("/texts/:textId", handler.DeleteText)

				protected.GET("/chapters", handler.ListChapters)
				protected.POST("/chapters", handler.CreateChapter)
				protected.GET("/chapters/:chapterId", handler.GetChapter)
				protected.PUT("/chapters/:chapterId", handler.UpdateChapter)
				protected.DELETE("/chapters/:chapterId", handler.DeleteChapter)

				protected.GET("/feedback", handler.ListFeedback)
				protected.POST("/feedback", handler.CreateFeedback)
				protected.GET("/feedback/:feedbackId", handler.GetFeedback)
				protected.PUT("/feedback/:feedbackId", handler.UpdateFeedback)
				protected.DELETE("/feedback/:feedbackId", handler.DeleteFeedback)

				protected.POST("/generate-prompt", handler.GeneratePrompt)

				// Ollama specific routes
				protected.POST("/ollama/generate", handler.GenerateResponse)
				protected.POST("/ollama/chat", handler.ChatWithModel)
			}
		} else {
			// Routes without authentication
			v1.GET("/users", handler.ListUsers)
			v1.POST("/users", handler.CreateUser)
			v1.GET("/users/:userId", handler.GetUser)
			v1.PUT("/users/:userId", handler.UpdateUser)
			v1.DELETE("/users/:userId", handler.DeleteUser)

			v1.GET("/trait-types", handler.ListTraitTypes)

			v1.GET("/texts", handler.ListTexts)
			v1.POST("/texts", handler.CreateText)
			v1.GET("/texts/:textId", handler.GetText)
			v1.PUT("/texts/:textId", handler.UpdateText)
			v1.DELETE("/texts/:textId", handler.DeleteText)

			v1.GET("/chapters", handler.ListChapters)
			v1.POST("/chapters", handler.CreateChapter)
			v1.GET("/chapters/:chapterId", handler.GetChapter)
			v1.PUT("/chapters/:chapterId", handler.UpdateChapter)
			v1.DELETE("/chapters/:chapterId", handler.DeleteChapter)

			v1.GET("/feedback", handler.ListFeedback)
			v1.POST("/feedback", handler.CreateFeedback)
			v1.GET("/feedback/:feedbackId", handler.GetFeedback)
			v1.PUT("/feedback/:feedbackId", handler.UpdateFeedback)
			v1.DELETE("/feedback/:feedbackId", handler.DeleteFeedback)

			v1.POST("/generate-prompt", handler.GeneratePrompt)

			// Ollama specific routes
			v1.POST("/ollama/generate", handler.GenerateResponse)
			v1.POST("/ollama/chat", handler.ChatWithModel)
		}
	}

	return router
}
