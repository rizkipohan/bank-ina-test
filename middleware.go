package main

import (
	"bank-ina-test/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		accessToken := c.GetHeader("Authorization")
		if accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized"})
			c.Abort()
			return
		}

		_, err := repository.ValidateToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// c.Set("userinfo", anyy.(repository.MyCustomClaims))

		// log.Println(c.Get("userinfo"))

		c.Next()

	}
}
