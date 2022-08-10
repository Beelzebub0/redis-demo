package lib

import (
	rdb "github.com/go-redis/redis/v8"
)

func GetRedisConn() *rdb.Client {
	return rdb.NewClient(&rdb.Options{
		Addr:     "localhost:6379",
		Password: "",
	})
}
