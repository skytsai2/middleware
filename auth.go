package router

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Printf("gogo %v \n", t)
	}
}
