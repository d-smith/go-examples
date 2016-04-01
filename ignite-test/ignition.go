package main


import (
	"github.com/alecthomas/kingpin"
	"os"
	"fmt"
	"math/rand"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"time"
	"errors"
	"sync"
	"sync/atomic"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

//command line options
var (
	app = kingpin.New("ignition", "Test program for an ignite cluster")
	verbose = app.Flag("verbose", "Enable verbose output").Bool()
	servers = app.Flag("server","hostname:port to send request to").Required().Strings()
	reads = app.Flag("reads", "Number of reads per write op").Default("1").Int()
	concurrent = app.Flag("concurrent","Number of go routines to execute write and reads").Default("1").Int()
	writes = app.Flag("writes","Total number of writes per go routine").Default("1").Int()
	waitTimeMillis = app.Flag("wait-time-millis", "Wait time between operations in ms").Default("500").Int()

)

//counters
var (
	writeCount uint64 = 0
	readCount uint64 = 0
	errorCount uint64 = 0
)

type IgniteResponse struct {
	AffinityNodeId string `json:"affinityNodeId"`
	Error string `json:"error"`
	Response interface{} `json:"response"`
	SessionToken string `json:"sessionToken"`
	SuccessStatus int `json:"sessionStatus"`
}

func handleErr(err error) {
	atomic.AddUint64(&errorCount, 1)
	log.Println(err.Error())
}

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	fmt.Println(*servers)

	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < *concurrent; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < *writes; j++ {
				writeAndRead(*reads, *waitTimeMillis, *verbose, *servers)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	duration := time.Now().Sub(start)

	log.Printf("%d writes, %d reads, %d errors in %v",
		writeCount, readCount, errorCount, duration)

}

func pickRandomServer(servers []string) string {
	rn := rand.Intn(MaxInt)
	idx := rn % len(servers)
	return servers[idx]
}

func generateID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return u.String(), nil
}


func writeAndRead(numReads, waitTimeMillis int, verbose bool, servers []string) {
	key,err := generateID()
	if err != nil {
		log.Fatal(err)
	}

	val := rand.Intn(MaxInt)
	endpoint := pickRandomServer(servers)
	if verbose {
		log.Printf("write %s->%d to %s\n",key,val,endpoint)
	}

	err = putItem(endpoint, key, val)
	if err != nil {
		handleErr(err)
		return
	}

	for i:= 0; i < numReads; i++ {
		time.Sleep(time.Duration(waitTimeMillis) * time.Millisecond)
		endpoint = pickRandomServer(servers)
		if verbose {
			log.Println("read from",endpoint)
		}
		readValue, err := getItem(endpoint,key)
		if err != nil {
			handleErr(err)
			return
		}

		log.Printf("read %d from %s for %s, wanted %d\n",
			readValue, endpoint,key,val)

		if readValue != val {
			handleErr(err)
			return
		}
	}



}

func putItem(endpoint, key string, val int) error {
	queryString :=fmt.Sprintf("http://%s/ignite?cmd=put&key=%s&val=%d&cacheName=xtSessionCache", endpoint, url.QueryEscape(key), val)
	resp, err := http.Get(queryString)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	_,err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	atomic.AddUint64(&writeCount, 1)
	return nil
}

func getItem(endpoint,key string)(int,error) {
	queryString :=fmt.Sprintf("http://%s/ignite?cmd=get&key=%s&cacheName=xtSessionCache", endpoint, url.QueryEscape(key))
	resp, err := http.Get(queryString)
	if err != nil {
		return -1,err
	}

	defer resp.Body.Close()
	saidTheServer,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1,err
	}

	var response IgniteResponse
	err = json.Unmarshal(saidTheServer,&response)
	if err != nil {
		return -1,err
	}

	responseStr,ok := (response.Response).(string)
	if !ok {
		return -1, errors.New("Unable to convert response to int")
	}

	intVal, err := strconv.Atoi(responseStr)
	if err != nil {
		return -1,err
	}

	atomic.AddUint64(&readCount, 1)
	return intVal,nil
}