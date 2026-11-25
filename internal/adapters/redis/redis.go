package redis

import (
	"context"
	"log"
	"strings"

	goredis "github.com/redis/go-redis/v9"
)

func NewClient(addr string) *goredis.Client {
	var opts *goredis.Options

	if strings.HasPrefix(addr, "redis://") || strings.HasPrefix(addr, "rediss://") {
		parsedOpts, err := goredis.ParseURL(addr)
		if err != nil {
			log.Fatalf("invalid redis URL: %v", err)
		}
		opts = parsedOpts
	} else {
		opts = &goredis.Options{
			Addr: addr,
		}
	}

	rdb := goredis.NewClient(opts)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("redis ping: %v", err)
	}
	log.Printf("connected to redis at %s", addr)

	return rdb
}
