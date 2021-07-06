package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "10.13.29.152:3606",
	})
	var c context.Context
	c, _ = context.WithCancel(context.Background())

	result, err := rdb.Ping(c).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}
