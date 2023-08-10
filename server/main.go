package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

const (
	ACCESS_LIMIT = 5
	COOLDOWN     = 2
)

var RC *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func main() {

	RC.FlushDB()

	r := gin.Default()
	r.GET("/buy", buy)
	r.Run("localhost:19810")
}

func buy(c *gin.Context) {
	cIP := c.ClientIP()
	freq, _ := RC.Incr(cIP).Result()
	RC.Expire(cIP, COOLDOWN*time.Second)

	if freq > ACCESS_LIMIT {
		c.JSON(429, gin.H{
			"result": "too many request",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": "buy item",
		"freq":   freq,
	})
}
