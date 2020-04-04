package middleware

import (
	"net/http"

	"app/kit"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		logrus.Infof("Get token: %s", token)

		auth := kit.NewAuthJWT()
		claims, err := auth.ParseToken(token)
		if err != nil {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
