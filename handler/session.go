package handler

import (
	"net/http"
	"time"

	"app/kit"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	auth := kit.NewAuthJWT()
	claims := &kit.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + 3600,
		},
		ID: "12345678",
	}

	token, err := auth.CreateToken(claims)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.SetCookie("token", token, 3600, "/", "", 1, false, true)
	c.JSON(200, gin.H{
		"message": "login success",
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", 1, false, true)
	c.JSON(200, gin.H{
		"message": "logout success",
	})
}

func RefreshToken(c *gin.Context) {
	claims := c.MustGet("claims").(*kit.AuthClaims)
	claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()

	auth := kit.NewAuthJWT()
	token, err := auth.CreateToken(claims)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.SetCookie("token", token, 3600, "/", "", 1, false, true)
	c.JSON(200, gin.H{
		"message": "refresh success",
	})
}
