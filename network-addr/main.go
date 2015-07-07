package main

import (
	"fmt"
	"net"
	"os"
)

func printHostName() {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Hostname: ", host)
}

func printInterfaceAddresses() {

	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Interface addresses:\n")
	for _, address := range addresses {
		fmt.Printf("\tnetwork form: %s, string form: %s\n", address.Network(), address.String())
	}
}

func printInterfaces() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Interfaces:")
	for _, i := range interfaces {
		fmt.Printf("\t%v\n", i)
	}
}

func printHostLookup() {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Addresses for host %s\n", host)
	for _, addr := range addrs {
		fmt.Printf("\t%s\n", addr)
	}
}

func main() {
	printHostName()
	printHostLookup()
	printInterfaceAddresses()
	printInterfaces()

}
