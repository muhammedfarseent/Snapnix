package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	controller "ECOMERS/Src/controller/user"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	r.POST("/signup", controller.UserSignUp)
	r.POST("/login", controller.Login)
	r.POST("/logout", controller.Logout)

	return r
}
