package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

type Object struct {
	Str string
	Num int
}

func main() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6380",
		},
	})
	mycache := cache.New(&cache.Options{
		Redis: ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	ctx := context.TODO()
	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}

	if err := mycache.Set(&cache.Item{
		Ctx: ctx,
		Key: key,
		Value: obj,
		TTL: time.Hour,
	}); err != nil {
		panic(err)
	}

	var wanted Object
	if err := mycache.Get(ctx, "mykey", &wanted); err == nil {
		fmt.Println(wanted)
	}
}