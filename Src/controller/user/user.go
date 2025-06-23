package controlller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	// "github.com/go-playground/validator"

	services "ECOMERS/Src/services/user"
	"ECOMERS/utils/models"
	response "ECOMERS/utils/response"
)

func UserSignUp(c *gin.Context) {
	var req models.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid request", nil, err.Error()))
		return
	}
	if err := validator.New().Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "validation failed", nil, err.Error()))
	}

	err := services.UserSignUp(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "signUp failed", req, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "signup succesfully otp sent", req, nil))
}

func Login(c *gin.Context) {
	var req models.LogInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid login data", nil, nil))
		return
	}
	if err := validator.New().Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "validation failed", nil, err.Error()))
		return
	}

	token, err := services.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "loggin failed ", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "Loggin succesfully", gin.H{"token": token}, nil))
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"messege":     "Logout succesfully",
	})
}
