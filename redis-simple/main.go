package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func setAndGet(client *redis.Client) error {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		return err
	}

	val, err := client.Get("key").Result()
	if err != nil {
		return err
	}
	fmt.Println("key", val)

	return nil
}

func otherCommonFunctions(client *redis.Client) error {
	err := client.Set("connection", 10, 0).Err()
	if err != nil {
		return err
	}

	intCmd := client.Incr("connection")
	if(intCmd.Err() != nil) {
		return intCmd.Err()
	}

	intCmd = client.Incr("connection")
	fmt.Println("two increments of connections yields", intCmd.Val())

	client.Del("connection")

	intCmd = client.Incr("connection")
	fmt.Println("increment of connection after del yields", intCmd.Val())

	return nil
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)

	setAndGet(client)
	otherCommonFunctions(client)
}
