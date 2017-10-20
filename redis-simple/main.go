package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
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

func lists(client *redis.Client)  {
	const key = "friends"
	client.RPush(key, "Alice")
	client.RPush(key, "Bob")
	client.LPush(key, "Sam")

	//Get all
	c := client.LRange(key, 0, -1)
	fmt.Println(c.Val())

	fmt.Println("length", client.LLen(key).Val())

	client.LPop(key)
	client.RPop(key)

	fmt.Println("post-pop length", client.LLen(key).Val())
}

func sets(client *redis.Client)  {
	const key = "superpowers"
	client.SAdd(key, "flight", "x-ray vision", "reflexes")
	fmt.Println(key, client.SMembers(key).Val())

	fmt.Println("set has flight?", client.SIsMember(key, "flight").Val())
	fmt.Println("set has focus?", client.SIsMember(key, "focus").Val())

	client.SAdd("bird powers", "flight", "song")

	fmt.Println("union of bird powers and super powers", client.SUnion("bird powers", key).Val())
}

func sortedSets(client *redis.Client)  {
	const key = "hackers"
	client.ZAdd(key, redis.Z{1940, "Alan Kay"})
	client.ZAdd(key, redis.Z{1906, "Grace Hopper"})
	client.ZAdd(key, redis.Z{1953, "Richard Stallman"})
	client.ZAdd(key, redis.Z{1965, "Yukihiro Matsumoto"})
	client.ZAdd(key, redis.Z{1916, "Claude Shannon"})
	client.ZAdd(key, redis.Z{1969, "Linus Torvalds"})
	client.ZAdd(key, redis.Z{1957, "Sophie Wilson"})
	client.ZAdd(key, redis.Z{1912, "Alan Turing"})

	fmt.Println("zrange", client.ZRange(key, 2, 4).Val())
}

func hashes(client *redis.Client)  {
	client.HSet("user:1000", "name", "john smith")
	client.HSet("user:1000", "email", "john.smith@example.com")
	client.HSet("user:1000", "password", "s3cret")

	fmt.Println("user:1000", client.HGetAll("user:1000").Val())
}

func hashFieldNumbers(client *redis.Client)  {
	client.HSet("user:1000", "visits", 10)
	client.HIncrBy("user:1000", "visits", 1)
	client.HIncrBy("user:1000", "visits", 10)

	fmt.Println("visits", client.HGet("user:1000", "visits").Val())

	client.HDel("user:1000", "visits")
	client.HIncrBy("user:1000", "visits", 1)
	fmt.Println("visits - inc after delete", client.HGet("user:1000", "visits").Val())

}

func keyTTLandExpiration(client *redis.Client) error {
	const key = "resource:lock"

	err := client.Set(key, "Redis demo", 0).Err()
	if err != nil {
		return err
	}

	boolCmd := client.Expire(key, 120*time.Second)
	if boolCmd.Err() != nil {
		return boolCmd.Err()
	}

	time.Sleep(2*time.Second)

	tc := client.TTL(key)
	fmt.Println("ttl", tc.Val())

	client.Set(key, "Redis demo 2",0)

	tc = client.TTL(key)
	fmt.Println("update key ttl", tc.Val())

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
	keyTTLandExpiration(client)
	lists(client)
	sets(client)
	sortedSets(client)
	hashes(client)
	hashFieldNumbers(client)
}
