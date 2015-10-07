package main

import (
	"crypto/x509"
	"encoding/pem"
	"log"
)

const certPEM = `
-----BEGIN CERTIFICATE-----
MIIC+DCCAeKgAwIBAgIRAIJaB8pAErenO9pMBUDo3awwCwYJKoZIhvcNAQELMBIx
EDAOBgNVBAoTB0FjbWUgQ28wHhcNMTUwODI4MTM0MzU3WhcNMTYwODI3MTM0MzU3
WjASMRAwDgYDVQQKEwdBY21lIENvMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEAwoH4nc/B3/i1D1TjCx3kgC6ygX3WHDv/xHtAoRAgHFUVElo3PznbxLAk
MvElVdAevCJaiuJaLiZARKLvwSJh08/9y+WMYa1nDjINk6UqG3huPXdJmTguzleO
c7UrCW4WKSo2HbeqYlF4BOiqnQhdDncUh5BgR8JXuiueMn2Ka59lkB/i+ryOt5W7
kaFKJhQEV67+fuES/5WfE+B4XsfT/ctXnGY0zrEInbJlyKwAzyCWJOJFrZte8cxs
235q3VMAhMRDU1IGNuWBIntfEXZgUXqI1Z9gsdbfTsQQ+xWhQCCOJwDrxAEg1Udk
dWn6NGWevsH4JoM9JzzOeSH8ZYPrVQIDAQABo00wSzAOBgNVHQ8BAf8EBAMCAKAw
EwYDVR0lBAwwCgYIKwYBBQUHAwEwDAYDVR0TAQH/BAIwADAWBgNVHREEDzANggtN
QUNMQjAxNTgwMzALBgkqhkiG9w0BAQsDggEBALJtJGaXx9At98CvEWKBpiGYqjUu
aiQHS5R61R/g8iqWkct77cqN6SBWTf138NZ3j3mvfROCoU96BEMEl0Fk9apLrikI
9Ns9/sl4nL1IOR56vddm46DfEV5CpMCAgrMGhFMJiaW4t9HvYjpBSs8T5n4tGqu/
JsvPhLGOcu5i4RiPpwM8f4fhnD3sija334jj5meJwg0NR8eO3ro1zaH+0hMQ7l8Q
tFJusSJenG28q9MXpOoCG6KLCmSCrIfDRYIpJQ0d5fXLO4YG92KFFqrf2ycOTydY
hN9G5ZWaErEY5j+sbYmeJBtEM5v6BQJotJh2SAh8RpYr69qJPLw6fdTu+mU=
-----END CERTIFICATE-----`

func main() {
	log.Println("decode cert PEM")
	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		log.Fatal("Unable to decode certificate PEM")
	}

	log.Println("parse the cert")
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal("failed to parse certificate: " + err.Error())
	}

	pk := cert.PublicKey

	pkbytes, err := x509.MarshalPKIXPublicKey(pk)
	if err != nil {
		log.Fatal("unable to marshal public key")
	}

	pemdata := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pkbytes,
		},
	)

	log.Println(string(pemdata))
}
