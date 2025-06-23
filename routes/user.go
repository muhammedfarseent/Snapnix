package routes

import (
	controller "SNAPNIX/Src/controller/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	r.POST("/signup", controller.UserSignUp)
	r.POST("/login", controller.Login)
	r.POST("/logout", controller.Logout)

	return r
}
