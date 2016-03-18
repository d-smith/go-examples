package main

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
	"os"
)

type Stuff struct {
	A string
	B string
	C int
}

func makeAndCacheStuff(mc *memcache.Client, stuffId string) {
	s := Stuff{
		A: "some A stuff",
		B: "some B stuff",
		C: 123,
	}

	valbytes, err := json.Marshal(&s)
	if err != nil {
		log.Fatal(err)
	}

	it := &memcache.Item{
		Key:        stuffId,
		Value:      valbytes,
		Expiration: 60,
	}

	err = mc.Set(it)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	//Create a client.
	mc := memcache.New("localhost:12000")

	it, err := mc.Get("stuff1")
	if err != nil {
		if err == memcache.ErrCacheMiss {
			log.Println("cache miss")
			makeAndCacheStuff(mc, "stuff1")
			os.Exit(0)
		}

		log.Fatal(err)
	}

	log.Println("Read item from the cache...")
	var stuff Stuff
	err = json.Unmarshal(it.Value, &stuff)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", stuff)

}
