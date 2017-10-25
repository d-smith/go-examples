package main

import (
	"fmt"
	"crypto/rand"
	"github.com/go-redis/redis"
	"time"
	"errors"
)

func uuid() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

type RateLimiter struct {
	intervalInMillis int64
	maxInInterval int
	client *redis.Client
}

func NewRateLimiter(intervalInMillis int64, maxInInterval int, client *redis.Client) *RateLimiter {
	return &RateLimiter{
		intervalInMillis:intervalInMillis,
		maxInInterval:maxInInterval,
		client:client,
	}
}

var concurrent_requests_limiter_lua = `
local key = KEYS[1]

local capacity = tonumber(ARGV[1])
local timestamp = tonumber(ARGV[2])
local id = ARGV[3]

local count = redis.call("zcard", key)
local allowed = count < capacity

if allowed then
	redis.call("zadd", key, timestamp, id)
end

return { allowed, count }
`

func (rl *RateLimiter) allowRequest(id string) (bool,error) {
	now := time.Now().UnixNano()/1000 //microseconds
	fmt.Println("now", now)
	clearBefore := now - (rl.intervalInMillis*1000) //microseconds
	fmt.Println("clearBefore", clearBefore)


	element,_ := uuid()
	fmt.Println("new element", element)

	rl.client.ZRemRangeByScore(id,"0",fmt.Sprintf("%d",clearBefore))
	defer func() {
		fmt.Println("keep the zset alive")
		boolCmd := rl.client.Expire(id, time.Duration(rl.intervalInMillis/1000)* time.Second)
		if boolCmd.Err() != nil {
			fmt.Println("warning - error setting expire on zset", boolCmd.Err().Error())
		}

	}()

	cmd := rl.client.Eval(concurrent_requests_limiter_lua, []string{id},rl.maxInInterval, now,element)
	if cmd.Err() != nil {
		fmt.Println("script execution error", cmd.Err().Error())
		return false, cmd.Err()
	}

	cmdOutput := cmd.Val()
//	fmt.Println(reflect.TypeOf(cmdOutput).String())
	outputSlice, ok := cmdOutput.([]interface{})
	if !ok {
		return false, errors.New("Unexcepted result type from Redis script execution")
	}

	return outputSlice[0] != nil, nil
//	fmt.Println(outputSlice)
//	if outputSlice[0] != nil {
//		fmt.Println("output 1 type", reflect.TypeOf(outputSlice[0]).String())
//	}


//	pipeline := rl.client.TxPipeline()
//	pipeline.ZRemRangeByScore(id,"0",fmt.Sprintf("%d",clearBefore))
//	pipeline.ZRangeWithScores(id, 0, -1)
//	pipeline.ZAdd(id,redis.Z{float64(now), element})

//	return true,nil
}



func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rl := NewRateLimiter(60000,6, client)
	allowed, err := rl.allowRequest("a1")
	if err != nil {
		fmt.Println("sigh", err.Error())
		return
	}

	fmt.Println("allowed?", allowed)
}
