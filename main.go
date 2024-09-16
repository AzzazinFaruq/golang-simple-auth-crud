package main

import (
	"Azzazin/backend/controllers"
	middleware "Azzazin/backend/middlewares"
	"Azzazin/backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//Declare New Gin Route System
	router := gin.New()
	//Run Database Setup
	models.ConnectDatabase()
	
	//For route that doesnt need a middleware like login, register, etc.
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	//for route that need auth with middleware
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/user", controllers.GetCurrentUser)
	protected.GET("/index", controllers.Index)
	protected.GET("/data/:id", controllers.ByID)
	protected.POST("/inputData", controllers.Input)
	protected.DELETE("/deleteData/:id", controllers.Delete)
	protected.PUT("/ubah/:id", controllers.Update)
	protected.POST("/logout", controllers.Logout)
    
	router.Run(":8000")
}