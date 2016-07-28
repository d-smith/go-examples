package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
	"log"
	"net"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

const (
	SaltBytes = 32
	UidBytes  = 32
)

func nonSaltyHash(mac string) string {
	hash := sha256.Sum256([]byte(mac))
	return fmt.Sprintf("%x", hash)
}

func saltyHash(mac string) string {
	salt := make([]byte, SaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	fatal(err)

	hash, err := scrypt.Key([]byte(mac), salt, 16384, 8, 1, UidBytes)
	fatal(err)

	return fmt.Sprintf("%x", hash)
}

func main() {
	ifaces, err := net.Interfaces()
	fatal(err)

	for _, iface := range ifaces {
		log.Printf("iface %s has mac address %s and its salty has is %s\n",
			iface.Name, iface.HardwareAddr, nonSaltyHash(iface.HardwareAddr.String()))
	}
}
