package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/xmlpath.v1"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

type Response struct {
	Name string  `json:name"Name"`
	Last float64 `json:name"Last"`
	Time string  `json:name"Time"`
	Date string  `json:name"Date"`
}

type QuoteRequest struct {
	Symbol string
}

var soapStart string = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://www.webserviceX.NET/">
<soapenv:Header/><soapenv:Body><web:GetQuote><web:symbol>`

var soapEnd string = `</web:symbol></web:GetQuote></soapenv:Body></soapenv:Envelope>`

func callQuoteService(symbol string) (string, error) {
	log.Println("call quote service for ", symbol)

	payload := fmt.Sprintf("%s%s%s", soapStart, symbol, soapEnd)

	client := &http.Client{}
	quoteReq, err := http.NewRequest("POST", "http://www.webservicex.net/stockquote.asmx", strings.NewReader(payload))
	if err != nil {
		return "", err
	}

	quoteReq.Header.Add("Content-Type", "text/xml")

	resp, err := client.Do(quoteReq)
	if err != nil {
		return "", err
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	resp.Body.Close()

	log.Println("quote service returned ", string(respData))
	return string(respData), nil
}

func quoteViaGet(r *http.Request) (string, error) {

	//Get the symbol as the last part of the request URI
	parts := strings.Split(r.RequestURI, "/")
	symbol := parts[len(parts)-1]
	fmt.Printf("requesting quote for %s\n", symbol)

	return callQuoteService(symbol)
}

func quoteViaPost(r *http.Request) (string, error) {
	//Parse the payload
	quoteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read post body")
		return "", nil
	}

	r.Body.Close()

	log.Println("Unmarshal quote body: ", string(quoteBody))
	var quote QuoteRequest
	err = json.Unmarshal(quoteBody, &quote)
	if err != nil {
		return "", err
	}

	return callQuoteService(quote.Symbol)
}

func getQuoteResponse(soapResponse string) (string, error) {
	compiledPath := xmlpath.MustCompile("/Envelope/Body/GetQuoteResponse/GetQuoteResult")
	root, err := xmlpath.Parse(strings.NewReader(soapResponse))
	if err != nil {
		return "", err
	}

	value, ok := compiledPath.String(root)
	if !ok {
		return "", errors.New("Unable to extract GetQuoteResult")
	}

	quoteResult := string(value)

	return quoteResult, nil
}

func getQuotePart(quoteResult string, partName string) (string, error) {
	compiledPath := xmlpath.MustCompile("/StockQuotes/Stock/" + partName)
	root, err := xmlpath.Parse(strings.NewReader(quoteResult))
	if err != nil {
		log.Fatal(err)
	}

	value, ok := compiledPath.String(root)
	if !ok {
		return "", errors.New("Unable to extract " + partName)
	}

	return string(value), nil
}

func formResponseJson(data string) ([]byte, error) {

	//Pull out the quote result part of the soap envelope
	quoteResult, err := getQuoteResponse(data)
	if err != nil {
		return nil, err
	}

	//From the quote result, pull out the name
	name, err := getQuotePart(quoteResult, "Name")
	if err != nil {
		return nil, nil
	}

	//Grab the last value
	last, err := getQuotePart(quoteResult, "Last")
	if err != nil {
		return nil, nil
	}

	lastFloat, err := strconv.ParseFloat(last, 64)
	if err != nil {
		return nil, nil
	}

	//Time
	time, err := getQuotePart(quoteResult, "Time")
	if err != nil {
		return nil, nil
	}

	//Date
	date, err := getQuotePart(quoteResult, "Date")
	if err != nil {
		return nil, nil
	}

	response := Response{
		Name: name,
		Last: lastFloat,
		Time: time,
		Date: date,
	}

	return json.Marshal(response)
}

func handleQuoteCalls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var quoteData string
	var err error

	switch r.Method {
	case "GET":
		log.Println("Handling GET request")
		quoteData, err = quoteViaGet(r)
	case "POST":
		log.Println("Handling POST request")
		quoteData, err = quoteViaPost(r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	responseJson, err := formResponseJson(quoteData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(responseJson))
}

func corsWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, r)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(rec.Body.Bytes())
	})
}

func main() {
	//Original static content
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/stuff/", http.StripPrefix("/stuff/", fs))

	//Add some dynamc content
	http.HandleFunc("/quote/", handleQuoteCalls)

	//Add the swagger spec
	ss := http.FileServer(http.Dir("dist"))
	http.Handle("/apispec/", http.StripPrefix("/apispec/", corsWrapper(ss)))

	log.Println("Listening...")
	http.ListenAndServe(":9000", nil)
}
