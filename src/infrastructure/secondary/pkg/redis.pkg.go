package pkg

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
)

var client *redis.Client
var onceRedis sync.Once

func newClient() *redis.Client {
	onceRedis.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
			DB:   1,
		})
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			panic(fmt.Errorf("error connecting to Redis: %v", err))
		}
	})

	return client
}

func GetClientRedis() *redis.Client {
	if client == nil {
		return newClient()
	}

	return client
}

func FlushAllRegisters() {
	client := GetClientRedis()
	err := client.FlushAll(context.Background()).Err()
	if err != nil {
		fmt.Println("Error flushing all registers: ", err)
	}
	c := cron.New(cron.WithSeconds())
	_, err = c.AddFunc("0 30 1 * * *", func() {
		err := client.FlushAll(context.Background()).Err()
		fmt.Println("Flushing all registers")
		if err != nil {
			fmt.Println("Error flushing all registers: ", err)
		}
	})

	if err != nil {
		fmt.Println("Error adding cron job: ", err)
	}

	c.Start()
}
