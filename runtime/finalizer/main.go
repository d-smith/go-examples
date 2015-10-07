package main

import (
	"fmt"
	"runtime"
	"time"
)

type Foo struct {
	Name string
}

func finalizeFoo(f *Foo) {
	fmt.Println("finalizing foo with name ", f.Name)
}

func doSomeFooStuff(name string) {
	foo := NewFoo(name)
	fmt.Printf("%v\n", foo)
}

func NewFoo(name string) *Foo {
	foo := &Foo{
		Name: name,
	}
	runtime.SetFinalizer(foo, finalizeFoo)
	return foo
}

func main() {
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("foo %d", i)
		doSomeFooStuff(name)
		time.Sleep(1 * time.Second)
		runtime.GC()
	}

	fmt.Println("done")
}
