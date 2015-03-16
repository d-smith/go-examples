package main 

/*
#include <stdlib.h>
*/
import "C"
import "fmt"
import "time"

func main() {
	seedInput := time.Now().Nanosecond()
	C.srandom(C.uint(seedInput))
	fmt.Println(C.random())	
}

