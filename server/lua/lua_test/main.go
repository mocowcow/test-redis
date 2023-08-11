package main

import (
	"fmt"
	"test-redis/server/lua"

	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	res, err := lua.BuyItem.Run(rdb, nil, 2).Int()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
