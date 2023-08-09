package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/buy", handler)
	r.Run("localhost:19810")
}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "buy item",
	})
}
