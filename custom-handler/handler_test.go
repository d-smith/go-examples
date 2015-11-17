package main


import (
	"testing"
	"net/http/httptest"
	"github.com/d-smith/go-examples/custom-handler/rc"
	"golang.org/x/net/context"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
)

func TestWithXRequestID(t *testing.T) {
	h := &rc.ContextAdapter{
		Ctx: context.Background(),
		Handler: rc.Middleware(rc.ContextHandlerFunc(handler)),
	}


	ts := httptest.NewServer(h)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL,nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	req.Header.Set("X-Request-ID", "request-id")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	if !strings.Contains(string(body), "request-id") {
		t.Fail()
	}
}
