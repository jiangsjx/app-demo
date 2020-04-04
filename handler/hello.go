package handler

import (
	"app/kit"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	claims := c.MustGet("claims").(*kit.AuthClaims)

	c.JSON(200, gin.H{
		"message": "hello " + claims.ID,
	})
}
