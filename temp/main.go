package main 

import (
	"io/ioutil"
	"fmt"
)

func main() {
	name, err := ioutil.TempDir("/tmp", "ex")
	if err != nil {
		fmt.Printf("error creating temp dir: %s\n", err.Error())
		return
	}
	
	fmt.Printf("Created directory %s\n", name)
	
	file, err := ioutil.TempFile(name, "ff")
	if err != nil {
		fmt.Printf("error creating temp file: %s\n", err.Error())
		return
	}
	
	fmt.Printf("Created file %s\n", file.Name())
}

