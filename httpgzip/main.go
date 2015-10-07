package main

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleGzipRequest() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		gr, err := gzip.NewReader(r.Body)
		fatal(err)

		defer gr.Close()
		plaintext, err := ioutil.ReadAll(gr)
		fatal(err)

		log.Println("Received some data:", string(plaintext))

		w.Write([]byte("bytes processed"))
	})
}

func createListener() (net.Listener, string) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	fatal(err)

	addr := "http://" + ln.Addr().String()
	return ln, addr
}

func startServiceWithListener(ln net.Listener, handler http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/", handler)

	server := &http.Server{
		Addr:    ln.Addr().String(),
		Handler: mux,
	}

	go server.Serve(ln)
}

func main() {
	log.Println("create listener")
	ln, addr := createListener()
	startServiceWithListener(ln, handleGzipRequest())

	log.Println("write to pipe")
	pr, pw := io.Pipe()
	gw := gzip.NewWriter(pw)
	in := strings.NewReader("This is a string to write yeah")

	go func() {
		_, err := io.Copy(gw, in)
		fatal(err)
		gw.Close()
		pw.Close()
	}()

	log.Println("create request")
	req, err := http.NewRequest("POST", addr, pr)
	fatal(err)

	log.Println("read response")
	resp, err := http.DefaultClient.Do(req)
	fatal(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fatal(err)
	log.Println("Got", string(body))
}
