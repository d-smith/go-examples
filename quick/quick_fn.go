package main

import (
	"fmt"
)

func OddMultipleOfThree(n int) int {
	fmt.Printf("OddMultipleOfThree called with argument %d\n", n)
	return 9
}

func StringConcat(s1 string, s2 string) string {
	fmt.Printf("StringConcat called with %v and %v\n", s1, s2)
	return s1 + s2
}

type Foo struct {
	Baz int
	Bar string
}

func BazBar(foo *Foo) string {
	fmt.Printf("BazBar called with %v\n", foo)
	return fmt.Sprintf("%d%s", foo.Baz, foo.Bar)
}
