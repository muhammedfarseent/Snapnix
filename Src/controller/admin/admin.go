package controlller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "ECOMERS/Src/services/admin"
)

type AuthHandler struct {
	Services *services.AuthService
}

func (h *AuthHandler) SignupAdmin(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.Services.SignupAdmin(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "admin registered", "user": user})
}

func (h *AuthHandler) LoginAdmin(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	token, err := h.Services.LoginAdmin(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messege": "Login seccessfully", "token": token})
}
