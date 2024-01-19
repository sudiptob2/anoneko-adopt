package router

import (
	"github.com/gin-gonic/gin"
	"anoneko-adopt-api/src/api/controllers"
)

func Init() {
	router := NewRouter()
	router.Run()
}

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Define the group of routes
	resource := router.Group("/api")

	// Add middleware functions to the group if needed
	// resource.Use(middleware1, middleware2, ...)

	// Define the routes in the group
	resource.GET("/pets", controllers.GetMysFits)

	return router
}
