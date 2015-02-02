package main

import (
	"fmt"
	"github.com/d-smith/go-examples/stringutil"
)

func main() {
	greeting := "Hello, world"
	doubleReveresed := stringutil.Reverse(stringutil.Reverse(greeting))
	fmt.Printf("%s\n", doubleReveresed)
}
