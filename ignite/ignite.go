package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"net/url"
	"encoding/json"
)

const baseUrl = "http://localhost:8080/ignite"

type IgniteResponse struct {
	AffinityNodeId string `json:"affinityNodeId"`
	Error string `json:"error"`
	Response interface{} `json:"response"`
	SessionToken string `json:"sessionToken"`
	SuccessStatus int `json:"sessionStatus"`
}

func putItem(key, value string) (*IgniteResponse,error) {
	queryString :=fmt.Sprintf("%s?cmd=put&key=%s&val=%s", baseUrl, url.QueryEscape(key), url.QueryEscape(value))
	println(queryString)
	resp, err := http.Get(queryString)
	if err != nil {
		return nil,err
	}

	defer resp.Body.Close()
	saidTheServer,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}

	var response IgniteResponse
	err = json.Unmarshal(saidTheServer,&response)
	if err != nil {
		return nil,err
	}

	return &response,nil
}

func getItem(key string)(*IgniteResponse,error) {
	queryString :=fmt.Sprintf("%s?cmd=get&key=%s", baseUrl, url.QueryEscape(key))
	println(queryString)
	resp, err := http.Get(queryString)
	if err != nil {
		return nil,err
	}

	defer resp.Body.Close()
	saidTheServer,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}

	var response IgniteResponse
	err = json.Unmarshal(saidTheServer,&response)
	if err != nil {
		return nil,err
	}

	return &response,nil
}

func main() {
	response,err := putItem("key x", "x xx xxx xx x")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)

	response,err = getItem("key x")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
