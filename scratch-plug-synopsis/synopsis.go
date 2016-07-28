package main

import (
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: synopsis <path to file containing synopsis>")
		os.Exit(1)
	}

	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		println("error opening synopsis file (", os.Args[1], "):", err.Error())
		os.Exit(1)
	}

	println(string(buf))
}
