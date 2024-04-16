package api

import (
	"github.com/danielsobrado/ainovelprompter/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	handler := &Handler{DB: db}

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		// Public routes
		v1.POST("/users/login", handler.Login)
		v1.POST("/users/register", handler.Register)

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/users", handler.ListUsers)
			protected.POST("/users", handler.CreateUser)
			protected.GET("/users/:userId", handler.GetUser)
			protected.PUT("/users/:userId", handler.UpdateUser)
			protected.DELETE("/users/:userId", handler.DeleteUser)

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
		}

		return router
	}
}
