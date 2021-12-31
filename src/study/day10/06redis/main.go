package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var redisDB *redis.Client

func initRedis() error {

	redisDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := redisDB.Ping(ctx).Result()
	return err

}

func main() {

	// 最新版本的go-redis库的相关命令都需要传递context.Context参数，例如：
	ctx := context.Background()
	err := initRedis()
	if err != nil {

		fmt.Println("redis链接错误", err.Error())
		return
	}
	fmt.Println("redis链接成功")

	{
		err := redisDB.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			panic(err)
		}
	}

	{
		val, err := redisDB.Get(ctx, "key").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("val:", val)
	}

}
