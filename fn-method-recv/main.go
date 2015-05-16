package main

import (
	"fmt"
)

type echoer func(string) string

func (fn echoer) decorate(prefix string, msg string) string {
	return fn(fmt.Sprintf("%s%s", prefix, msg))
}

func main() {
	var e echoer = func(s string) string {
		return s
	}

	fmt.Println(e("Hello"))
	fmt.Println(e.decorate("==>", "Hello"))
}
