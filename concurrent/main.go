package main

import "fmt"

func main() {
	fmt.Println("Synchronized verson:")
	SyncMain()
	
	fmt.Println("Message passing version:")
	MsgPassMain()
}

