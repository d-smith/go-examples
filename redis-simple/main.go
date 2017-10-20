package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	/*
	log.Print("Connect to redis...")
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected... deferring close.")
	defer conn.Close()

	log.Print("Add some keys.")
	if _, err = conn.Do("SET", "k1", "k1 value"); err != nil {
		log.Fatal(err)
	}
	if _, err = conn.Do("SET", "k2", "k2 value"); err != nil {
		log.Fatal(err)
	}

	log.Print("Read back the keys")
	strs, err := redis.Strings(conn.Do("MGET", "k1", "k2"))
	if err != nil {
		log.Fatal(err)
	}

	log.Print(strs)

	log.Print("Look up k1")

	now := time.Now()
	str, err := redis.String(conn.Do("GET", "k1"))
	duration := time.Since(now)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(str)
	log.Print("key lookup took duration ", duration)
	*/
}
