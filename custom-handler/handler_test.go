package main

import (
	"github.com/d-smith/go-examples/custom-handler/customctx"
	"github.com/d-smith/go-examples/custom-handler/customctx/reqid"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWithXRequestID(t *testing.T) {
	h := &customctx.ContextAdapter{
		Ctx:     context.Background(),
		Handler: reqid.Middleware(customctx.ContextHandlerFunc(handler)),
	}

	ts := httptest.NewServer(h)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
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
