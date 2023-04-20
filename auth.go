package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type checkHeader struct {
	Key string `header:"key" binding:"required"`
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := checkHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"errorCode": "01",
				"status":    "502",
			})
		}
		c.Next()
	}
}
