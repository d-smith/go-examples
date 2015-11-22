package main

import (
	"encoding/xml"
)

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
