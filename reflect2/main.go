package main
import (
	"reflect"
	"fmt"
)


var fooFactories map[string]FooFactory

func init() {
	fooFactories = make(map[string]FooFactory)
}

type Foo interface {
	FooThings(int) int
}

type AFoo struct{}

func (af AFoo) FooThings(x int) int {
	return x * 2
}

type BFoo struct{}

func(bf BFoo) FooThings(x int) int {
	return x * 3
}

type FooFactory func() Foo



func NewAFoo() Foo {
	return &AFoo{}
}

func NewBFoo() Foo {
	return &BFoo{}
}

func RegisterFactory(name string, factory FooFactory) {
	fooFactories[name] = factory
}


func GetFactory(name string) FooFactory {
	return fooFactories[name]
}


func main() {

	RegisterFactory("a", NewAFoo)
	RegisterFactory("b", NewBFoo)

	aFooFactory := GetFactory("a")

	theFoo := aFooFactory()

	typeOfFoo := reflect.TypeOf(theFoo)
	valueOfFoo := reflect.ValueOf(theFoo)

	fmt.Printf("Type of theFoo: %v\n", typeOfFoo)

	_,isFoo := valueOfFoo.Interface().(Foo)
	fmt.Println("Is Foo?", isFoo)

	_, isFooFactory := valueOfFoo.Interface().(FooFactory)
	fmt.Println("Is FooFactory?", isFooFactory)

}