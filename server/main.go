package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var RC *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

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
