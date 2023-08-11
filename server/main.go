package main

import (
	"fmt"
	"strconv"
	"test-redis/server/lua"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

const (
	ACCESS_LIMIT = 100
	COOLDOWN     = 2
)

var RC *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func main() {

	RC.FlushDB()
	initGoods()

	r := gin.Default()
	r.Use(acccessLimitMiddleware)

	r.GET("/buyRacing/:amount", buyWithRacing)
	r.GET("/buy/:amount", buy)
	r.Run("localhost:19810")
}

func initGoods() {
	RC.Set("goodsTotal", "20", 0)
	RC.Set("goodsSold", "0", 0)
}

func acccessLimitMiddleware(c *gin.Context) {
	cIP := c.ClientIP()
	freq, _ := RC.Incr(cIP).Result()
	RC.Expire(cIP, COOLDOWN*time.Second)

	if freq > ACCESS_LIMIT {
		c.JSON(429, gin.H{
			"result": "too many request",
		})
		c.Abort()
	}
}

func buyWithRacing(c *gin.Context) {
	// bug version
	// has race condition
	amountStr := c.Param("amount")
	amount, _ := strconv.Atoi(amountStr)
	totalStr, _ := RC.Get("goodsTotal").Result()
	total, _ := strconv.Atoi(totalStr)
	soldStr, _ := RC.Get("goodsSold").Result()
	sold, _ := strconv.Atoi(soldStr)

	if sold+amount > total {
		c.JSON(403, gin.H{
			"result": "insufficient stock",
		})
		return
	}

	RC.IncrBy("goodsSold", int64(amount))

	c.JSON(200, gin.H{
		"result": fmt.Sprintf("buy %d item", amount),
	})
}

func buy(c *gin.Context) {
	amountStr := c.Param("amount")
	amount, _ := strconv.Atoi(amountStr)

	res, err := lua.BuyItem.Run(RC, nil, amount).Int()

	if err != nil {
		c.JSON(403, gin.H{
			"result": err,
		})
		return
	}

	if res == 0 {
		c.JSON(403, gin.H{
			"result": "insufficient stock",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": fmt.Sprintf("buy %d item", amount),
	})
}
