package quote

import (
	"bytes"
	"encoding/xml"
	"fmt"
	cc "github.com/d-smith/go-examples/custom-handler/customctx"
	"github.com/d-smith/go-examples/custom-handler/customctx/timing"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

func extractResource(uri string) (string, error) {
	parts := strings.Split(uri, "/")
	if len(parts) != 3 || parts[2] == "" {
		return "", fmt.Errorf("Expected URI format: /quote/<symbol>")
	}

	return parts[2], nil

}

type quoteEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    body
}

type body struct {
	GetLastTradePrice lastTradePrice
}

type lastTradePrice struct {
	Symbol string
}

type responseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    responseBody
}

type responseBody struct {
	XMLName                   xml.Name `xml:"Body"`
	GetLastTradePriceResponse lastTradePriceResponse
}

type lastTradePriceResponse struct {
	Price string
}

func getQuoteRequestForSymbol(symbol string) quoteEnvelope {
	return quoteEnvelope{
		Body: body{
			GetLastTradePrice: lastTradePrice{Symbol: symbol},
		},
	}
}

var transport = &http.Transport{DisableKeepAlives: false, DisableCompression: false}

//NewQuoteHandler creates a quote handler with the given SOAP service endpoint
func NewQuoteHandler(hostAndPort string) func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {

		const timerName = "backend service call"
		timing.StartTimer(ctx, timerName)

		req.URL.Scheme = "http"
		req.URL.Host = hostAndPort
		req.Host = hostAndPort
		resp, err := transport.RoundTrip(req)
		if err != nil {
			println(err.Error())
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			timing.EndTimer(ctx, timerName, err)
			return
		}
		io.Copy(rw, resp.Body)
		resp.Body.Close()

		timing.EndTimer(ctx, timerName, nil)
	}
}

//Middleware returns a context aware wrapper that converts a GET on a stock symbol
//to a SOAP request to the quote service
func Middleware(ctxHandler cc.ContextHandler) cc.ContextHandler {
	return cc.ContextHandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		const timerName = "message and protocol transformation"
		timing.StartTimer(ctx, timerName)

		//Grab the symbol to quote from the uri
		resourceID, err := extractResource(r.RequestURI)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		println("quote for", resourceID)

		//Convert the method to POST for SOAP, and set the soap service
		//endpoint for the destination server
		r.Method = "POST"
		r.URL.Path = "/services/quote/getquote"

		//Form the SOAP payload
		payload := getQuoteRequestForSymbol(resourceID)
		payloadBytes, err := xml.Marshal(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Post the payload, and record the response
		r.Body = ioutil.NopCloser(bytes.NewReader(payloadBytes))
		rec := httptest.NewRecorder()

		ctxHandler.ServeHTTPContext(ctx, rec, r)

		//Parse the recorded response to allow the quote price to be extracted
		var response responseEnvelope
		err = xml.Unmarshal(rec.Body.Bytes(), &response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			timing.EndTimer(ctx, timerName, err)
			return
		}

		//Return just the price to the caller
		w.Write([]byte(response.Body.GetLastTradePriceResponse.Price + "\n"))
		timing.EndTimer(ctx, timerName, nil)

	})
}
