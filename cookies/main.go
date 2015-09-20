package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
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

func handleRequest() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookies := r.Cookies()
		fmt.Println("----request cookies---")
		for _, c := range cookies {
			printCookie(c)
		}
		fmt.Println("received ", len(r.Cookies()), " cookies")

		cookie := &http.Cookie{
			Name:  "MyCookie",
			Value: "Something to bring back",
		}

		http.SetCookie(w, cookie)
		w.Write([]byte("a response"))
	})
}

//Don't ever do this for a real application - it accepts cookies for any domain
//from any domain.
type TestSuffixList struct{}

func (ts *TestSuffixList) PublicSuffix(domain string) string {
	return ""
}

func (ts *TestSuffixList) String() string {
	return "for test purposes only"
}

func printCookie(cookie *http.Cookie) {
	fmt.Println("Cookie - name: ", cookie.Name, " value: ", cookie.Value)
}

func main() {
	ln, addr := createListener()
	startServiceWithListener(ln, handleRequest())

	options := cookiejar.Options{
		PublicSuffixList: new(TestSuffixList),
	}

	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{Jar: jar}

	resp, err := client.Get(addr + "/foo")
	fatal(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fatal(err)

	fmt.Println(string(body))
	url, _ := url.Parse(addr)
	fmt.Printf("client has %d cookies\n", len(client.Jar.Cookies(url)))

	resp, err = client.Get(addr + "/foo")
}
