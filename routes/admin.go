package routes

import (
	controller "SNAPNIX/Src/controller/admin"
	repository "SNAPNIX/Src/repository/admin"
	services "SNAPNIX/Src/services/admin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
