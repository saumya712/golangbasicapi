package routes

import (
	"GINPROJ/controllers"

	"github.com/gin-gonic/gin"
)

func Userroutes(r *gin.Engine) {
	r.POST("/users", controllers.Createuser)
	r.GET("/users", controllers.Getuser)
}
