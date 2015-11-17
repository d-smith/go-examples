package main

import (
	"encoding/xml"
)

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
