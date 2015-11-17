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

type QuoteEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body
}

type Body struct {
	GetLastTradePrice LastTradePrice
}

type LastTradePrice struct {
	Symbol string
}

type ResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    ResponseBody
}

type ResponseBody struct {
	XMLName                   xml.Name `xml:"Body"`
	GetLastTradePriceResponse LastTradePriceResponse
}

type LastTradePriceResponse struct {
	Price string
}

func getQuoteRequestForSymbol(symbol string) QuoteEnvelope {
	return QuoteEnvelope{
		Body: Body{
			GetLastTradePrice: LastTradePrice{Symbol: symbol},
		},
	}
}

var transport = &http.Transport{DisableKeepAlives: false, DisableCompression: false}

func QuoteHandler(ctx context.Context, rw http.ResponseWriter, req *http.Request) {

	const timerName = "backend service call"
	timing.StartTimer(ctx, timerName)

	req.URL.Scheme = "http"
	req.URL.Host = "localhost:4545"
	req.Host = "localhost:4545"
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

func QuoteMiddleware(ctxHandler cc.ContextHandler) cc.ContextHandler {
	return cc.ContextHandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {

		const timerName = "message and protocol transformation"
		timing.StartTimer(ctx, timerName)

		//Grab the symbol to quote from the uri
		resourceId, err := extractResource(r.RequestURI)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		println("quote for", resourceId)

		//Convert the method to POST for SOAP, and set the soap service
		//endpoint for the destination server
		r.Method = "POST"
		r.URL.Path = "/services/quote/getquote"

		//Form the SOAP payload
		payload := getQuoteRequestForSymbol(resourceId)
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
		var response ResponseEnvelope
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
