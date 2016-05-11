package main

import (
	"reflect"
	"fmt"
)

type Object interface{}

type Array []interface{}

type AThing struct {
	A string
	SomeOtherStuff string
}

func main() {
	a := AThing{
		A: "Hey I'm an A thing",
	}

	aType := reflect.TypeOf(a)
	fmt.Printf("%v with kind %s with %d field(s)\n", aType, aType.Kind(), aType.NumField())

	fmt.Println(reflect.TypeOf(new(Object)))

	var s string = aType.String()
	fmt.Println(s)
}