package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {

	t.Log("Given a server handling requests with echoHandler")
	{
		ts := httptest.NewServer(http.HandlerFunc(echoHandler))
		defer ts.Close()

		t.Log("Test server URL: ", ts.URL)
		t.Log("When going a get on /echo.test")
		{
			echoTestEndpoint := ts.URL + "/echo/test"
			res, err := http.Get(echoTestEndpoint)
			if err != nil {
				log.Fatal(err)
			}

			t.Log("Expect a response of test\\n")
			{
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Fatal(err)
				}

				if string(body) != "test\n" {
					log.Fatal("expected 'test' followed by a newline")
				}
			}
		}
	}
}

func TestEchoUsingAsserts(t *testing.T) {

	t.Log("Given a server handling requests with echoHandler")
	{
		ts := httptest.NewServer(http.HandlerFunc(echoHandler))
		defer ts.Close()

		t.Log("Test server URL: ", ts.URL)
		t.Log("When going a get on /echo.test")
		{
			echoTestEndpoint := ts.URL + "/echo/test"
			res, err := http.Get(echoTestEndpoint)
			assert.Nil(t, err)

			t.Log("Expect a response of test\\n")
			{
				body, err := ioutil.ReadAll(res.Body)
				assert.Nil(t, err)
				assert.Equal(t, "test\n", string(body), "Expect test\\n")
			}
		}
	}
}
