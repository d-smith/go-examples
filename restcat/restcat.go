package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	GET = "GET"
	PUT = "PUT"
)

var items map[string]Item

func init() {
	items = make(map[string]Item)
}

type Item struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/item/"):]
	switch r.Method {
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	case GET:
		getItem(id, r, w)
	case PUT:
		putItem(id, r, w)
	}
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	values := make([]Item, len(items))
	idx := 0
	for _, v := range items {
		values[idx] = v
		idx++
	}

	wireItems, err := json.Marshal(values)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(wireItems)
}

func getItem(id string, r *http.Request, w http.ResponseWriter) {

	item, ok := items[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	wireItem, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(wireItem)

}

func putItem(id string, r *http.Request, w http.ResponseWriter) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	r.Body.Close()

	item := Item{Id: id}

	err = json.Unmarshal(body, &item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	items[id] = item

}

func main() {
	http.HandleFunc("/item/", itemHandler)
	http.HandleFunc("/items", itemsHandler)
	http.ListenAndServe(":4000", nil)
}
