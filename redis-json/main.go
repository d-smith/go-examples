package main

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
)

type Foobar struct {
	Name string
	Foos []*Foo
	Bars []*Bar
}

type Foo struct {
	Name  string
	Level int
}

type Bar struct {
	Name  string
	Level int
}

func NewFoobar() *Foobar {
	return &Foobar{}
}

func (f *Foobar) AddFoo(foo *Foo) {
	f.Foos = append(f.Foos, foo)
}

func (f *Foobar) AddBar(bar *Bar) {
	f.Bars = append(f.Bars, bar)
}

func makeAFoobar(name string) *Foobar {
	fb := NewFoobar()
	fb.Name = name
	foo := &Foo{Name: "foo", Level: 50}
	bar1 := &Bar{Name: "bar1", Level: 15}
	bar2 := &Bar{Name: "bar3", Level: 19}
	fb.AddFoo(foo)
	fb.AddBar(bar1)
	fb.AddBar(bar2)

	return fb
}

func foosEqual(f1 *Foo, f2 *Foo) bool {
	return f1.Name == f2.Name && f1.Level == f2.Level
}

func barsEqual(b1 *Bar, b2 *Bar) bool {
	return b1.Name == b2.Name && b1.Level == b2.Level
}

func foobarsEqual(f1 *Foobar, f2 *Foobar) bool {
	if f1 == nil || f2 == nil {
		return f1 == f2
	}

	if f1.Name != f2.Name {
		return false
	}

	if len(f1.Foos) != len(f2.Foos) {
		return false
	}

	for i := 0; i < len(f1.Foos); i++ {
		if !foosEqual(f1.Foos[i], f2.Foos[i]) {
			return false
		}
	}

	if len(f1.Bars) != len(f2.Bars) {
		return false
	}

	for i := 0; i < len(f1.Bars); i++ {
		if !barsEqual(f1.Bars[i], f2.Bars[i]) {
			return false
		}
	}

	return true
}

func main() {
	log.Print("Connect to redis...")
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected... deferring close.")
	defer conn.Close()

	const fooName = "a foobar"

	foobar := makeAFoobar(fooName)

	jsonFB, _ := json.Marshal(foobar)
	log.Println("store: ", string(jsonFB))

	if _, err := conn.Do("SET", fooName, jsonFB); err != nil {
		log.Fatal(err)
	}

	log.Println("Read the json back as a string")
	fooJson, err := redis.String(conn.Do("GET", fooName))
	if err != nil {
		log.Fatal(err)
	}

	log.Print("read back ", fooJson)
	log.Print("now whip up another foo struct from the same json")

	var sonOfFoobar Foobar

	if err = json.Unmarshal([]byte(fooJson), &sonOfFoobar); err != nil {
		log.Fatal(err)
	}

	log.Println("compare before and after")
	if foobarsEqual(foobar, &sonOfFoobar) {
		log.Println("the two Foobars are the same")
	} else {
		log.Println("hmmmm... the two Foobars differ")
	}
}
