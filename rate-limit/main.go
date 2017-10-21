package main

import (
	"time"
	"github.com/go-redis/redis"
	"fmt"
	"crypto/rand"
	"reflect"
	"strings"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}


func GenerateUuidV4() (string, error) {
	//16 random bytes
	bytes, err := GenerateRandomBytes(16)
	if err != nil {
		return "", err
	}
	//set v4 byte
	bytes[6] = (bytes[6] & 0xf) | 0x4<<4
	//set version rfc4122
	bytes[8] = (bytes[8] & 0x3f) | 0x80

	return fmt.Sprintf("%x-%x-%x-%x-%x", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:]), nil
}

func uuid() string {
	genUUID,_ := GenerateUuidV4()
	return genUUID
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

func (rl *RateLimiter) timeLeft(id string) (int,error) {
	now := time.Now().UnixNano()/1e6
	fmt.Println("now", now)
	clearBefore := now - (rl.intervalInMillis*1000)
	fmt.Println("clearBefore", clearBefore)


	pipeline := rl.client.TxPipeline()
	pipeline.ZRemRangeByScore(id,"0",fmt.Sprintf("%d",clearBefore))
	pipeline.ZRangeWithScores(id, 0, -1)
	pipeline.ZAdd(id,redis.Z{float64(now), uuid()})
	pipeline.Expire(id, time.Duration(rl.intervalInMillis/1000)* time.Second)

	var cmdErr []redis.Cmder
	var pipelineErr error
	cmdErr, pipelineErr = pipeline.Exec()
	if pipelineErr != nil {
		return -1,pipelineErr
	}

	rangeWithScoresResult := cmdErr[1]
	fmt.Println("rwsr name",rangeWithScoresResult.Name())
	fmt.Println("rwsr args",rangeWithScoresResult.Args())
	fmt.Println("rwsr strings",rangeWithScoresResult.String())

	fmt.Println(reflect.TypeOf(rangeWithScoresResult.String()).String())
	parts := strings.Split(rangeWithScoresResult.String(), " ")
	for _,p := range parts {
		fmt.Println(p)
	}


	return 1,nil
}


func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rl := NewRateLimiter(60000,6, client)
	rl.timeLeft("a1")
}


