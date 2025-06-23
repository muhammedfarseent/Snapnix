package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	controller "ECOMERS/Src/controller/admin"
	repository "ECOMERS/Src/repository/admin"
	services "ECOMERS/Src/services/admin"
)

func AdminRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	//adminRoutes
	adminRepo := &repository.AdminRepository{DB: db}
	authService := &services.AuthService{Repo: adminRepo}
	authHandler := &controller.AuthHandler{Services: authService}

	adminAuth := r.Group("/Authentication")
	{
		adminAuth.POST("/signup", authHandler.SignupAdmin)
		adminAuth.POST("/login", authHandler.LoginAdmin)
	}
	return r

}
