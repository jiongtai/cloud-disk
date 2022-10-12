package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "",
	DB:       0,
})

func TestRedis(t *testing.T) {
	err := rdb.Set(ctx, "name", "tyson", time.Second*10).Err()
	if err != nil {
		t.Fatal(err)
	}
	result, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}
