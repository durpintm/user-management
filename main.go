package main

import (
	"github.com/durpintm/user-management/controllers"
	"github.com/durpintm/user-management/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)

	router.Run()
}
