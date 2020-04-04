package main

import (
	"app/handler"
	"app/middleware"

	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/api/v1/login", handler.Login)

	v1 := router.Group("/api/v1", middleware.JWTAuth())
	{
		v1.GET("/hello", handler.Hello)
		v1.GET("/logout", handler.Logout)
		v1.GET("/refresh", handler.RefreshToken)
	}

	return router
}
