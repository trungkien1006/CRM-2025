package middlewares

import (
	"admin-v1/app/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthGuard(c *gin.Context) {
	var jwt string = c.GetHeader("Authorization")

	if err := helpers.CheckJWT(jwt); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		c.Abort()

		return
	}	
	
	c.Next()
}