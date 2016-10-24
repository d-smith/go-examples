package main

import (
	"log"
	"github.com/hydrogen18/test-tls"
	"net/http"
	"os"
	"crypto/tls"
	"fmt"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Need three args: server key file, server cert file, ca cert file")
		os.Exit(1)
	}


	log.Println("get tls config")
	config := common.MustGetTlsConfiguration()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		w.Write([]byte("Here you go, chief"))
	})

	server := &http.Server{
		Addr:":51000",
		Handler:mux,
		TLSConfig:config,
		TLSNextProto:make(map[string]func(*http.Server,*tls.Conn,http.Handler), 0),
	}

	l := len(os.Args)
	privateKeyFile := os.Args[l-3]
	certificateFile := os.Args[l-2]
	log.Printf("Cert file is %s", certificateFile)
	log.Printf("Key file is %s", privateKeyFile)
	log.Fatal(server.ListenAndServeTLS(certificateFile, privateKeyFile))
}