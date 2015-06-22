package main

import (
	"fmt"
	"net"
)

func main() {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, address := range addresses {
		fmt.Printf("network form: %s, string form: %s\n", address.Network(), address.String())
	}

}
