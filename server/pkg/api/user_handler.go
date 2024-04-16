package api

import (
	"net/http"
	"strconv"

	"github.com/danielsobrado/ainovelprompter/pkg/auth"
	"github.com/danielsobrado/ainovelprompter/pkg/logging"
	"github.com/danielsobrado/ainovelprompter/pkg/models"
	"github.com/danielsobrado/ainovelprompter/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logging.Logger.Errorf("Failed to bind login input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var user models.User
	if err := h.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch user: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.HashedPassword) {
		logging.Logger.Error("Invalid password")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		logging.Logger.Errorf("Failed to generate token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logging.Logger.Errorf("Failed to bind registration input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		logging.Logger.Errorf("Failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	user := models.User{
		Username:       input.Username,
		Email:          input.Email,
		HashedPassword: hashedPassword,
		Role:           "user",
	}

	if err := h.DB.Create(&user).Error; err != nil {
		logging.Logger.Errorf("Failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *Handler) ListUsers(c *gin.Context) {
	search := c.Query("search")

	var users []models.User
	query := h.DB.Order("created_at DESC")
	if search != "" {
		query = query.Where("username ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Find(&users).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logging.Logger.Errorf("Failed to bind user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.DB.Create(&user).Error; err != nil {
		logging.Logger.Errorf("Failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) GetUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		logging.Logger.Errorf("Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch user: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		logging.Logger.Errorf("Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch user: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		logging.Logger.Errorf("Failed to bind user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.DB.Save(&user).Error; err != nil {
		logging.Logger.Errorf("Failed to update user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		logging.Logger.Errorf("Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		logging.Logger.Errorf("Failed to fetch user: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := h.DB.Delete(&user).Error; err != nil {
		logging.Logger.Errorf("Failed to delete user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.Status(http.StatusNoContent)
}
