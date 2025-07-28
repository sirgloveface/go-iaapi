package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirgloveface/go-iaapi/internal/auth"
)

func LoginHandler(c *gin.Context) {
	type LoginRequest struct {
		UserID string `json:"user_id"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user_id"})
		return
	}

	token, err := auth.GenerateJWT(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
