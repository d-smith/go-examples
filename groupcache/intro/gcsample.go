package main

import (
	"github.com/golang/groupcache"
	"log"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	me := "http://localhost:9000"
	groupcache.NewHTTPPool(me)

	var thingGroup = groupcache.NewGroup("thing-group", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			log.Println("thing group getter called with key", key)
			dest.SetString(key + " - value")
			return nil
		},
	))

	var sval string
	log.Println("get foo")
	err := thingGroup.Get(nil, "foo", groupcache.StringSink(&sval))
	fatal(err)

	log.Println("read foo val", sval)

	log.Println("get foo")
	err = thingGroup.Get(nil, "foo", groupcache.StringSink(&sval))
	fatal(err)

	log.Println("read foo val", sval)

	log.Println("get bar")
	err = thingGroup.Get(nil, "bar", groupcache.StringSink(&sval))
	fatal(err)

	log.Println("read bar val", sval)
}
