package main

import (
	"time"
	"github.com/go-redis/redis"
	"fmt"
	"crypto/rand"
	"strings"
	"strconv"
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
	now := time.Now().UnixNano()/1000 //microseconds
	fmt.Println("now", now)
	clearBefore := now - (rl.intervalInMillis*1000) //microseconds
	fmt.Println("clearBefore", clearBefore)


	element := uuid()
	fmt.Println("new element", element)

	pipeline := rl.client.TxPipeline()
	pipeline.ZRemRangeByScore(id,"0",fmt.Sprintf("%d",clearBefore))
	pipeline.ZRangeWithScores(id, 0, -1)
	pipeline.ZAdd(id,redis.Z{float64(now), element})
	pipeline.Expire(id, time.Duration(rl.intervalInMillis/1000)* time.Second)

	var cmdErr []redis.Cmder
	var pipelineErr error
	cmdErr, pipelineErr = pipeline.Exec()
	if pipelineErr != nil {
		return -1,pipelineErr
	}

	fmt.Println("rem range result", cmdErr[0])
	rangeWithScoresResult := cmdErr[1]

	elements := zparts(rangeWithScoresResult.String())
	fmt.Println(elements)

	//Ok to keep making requests?
	if len(elements) < rl.maxInInterval {
		return 0, nil
	}

	//Since the max requests for the interval have been made, the next time a request
	//can be made is when a slot opens up in the zparts. That will the difference between
	//the timestamp of the oldest element and the length of the interval.
	timeleft := (int64(elements[0].Score) - clearBefore)/1000  //divide by 1000 to return time left in ms

	fmt.Println("timeleft", timeleft)


	return 1,nil
}


//Looks like [{1.508595813639e+12 1a2942d7-f536-4b6b-a312-88c1465c18c5} {1.508595815287e+12 6bc8a7a7-1435-4a33-8b7d-32edabedd61b}]
func zparts(zstring string) []redis.Z {
	var elements []redis.Z

	parts := strings.Split(zstring, ":")

	if len(parts) < 2 {
		return elements
	}

	zslice := strings.TrimSpace(parts[1])
	if zslice == "[]" {
		return elements
	}

	zslice = strings.Trim(zslice,"[]")
	zpairs := strings.Split(zslice, " ")




	for i := 0; i < len(zpairs); i += 2 {
		rawScore := zpairs[i]
		rawElement := zpairs[i+1]

		score,_ := strconv.ParseFloat(strings.Trim(rawScore, "{"),64)
		element := strings.Trim(rawElement,"}")

		z := redis.Z{
			Score:score,
			Member:element,
		}

		elements = append(elements,z)
	}

	return elements
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


