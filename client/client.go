package client

import (
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
)

var (
	Rdb *redis.Client
)

func initClient() error {
	conn, err := grpc.Dial("")
	if err != nil {
		return err
	}
	defer conn.Close()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}
