package main

import (
	"sampleJWTRestAPI/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := setRouter()
	r.Run(":8000")
}

func setRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controller.RegisterUser())
	r.POST("/login", controller.LoginUser())

	auth := r.Group("/auth")
	auth.GET("/refreshToken", controller.RefreshToken())
	auth.Use(controller.AuthUser())
	{
		auth.GET("/userInfo", controller.GetUser())
		auth.POST("/updateUser", controller.UpdateUser())
	}

	return r
}
