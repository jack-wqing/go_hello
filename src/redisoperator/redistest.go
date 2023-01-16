package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	defer rdb.Close()
	statusCmd := rdb.Set(ctx, "go-redis-test", "test-value", 0)
	if statusCmd.Val() != "OK" {
		fmt.Println("rdb.Set statusCmd.val()=", statusCmd.Val())
	}

}
