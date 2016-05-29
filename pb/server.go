package main

import (
	"github.com/golang/protobuf/proto"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myClient := Client{}

		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Println(err)
		}

		if err := proto.Unmarshal(data, &myClient); err != nil {
			fmt.Println(err)
		}

		println(myClient.Id, ":", myClient.Name, ":", myClient.Email, ":", myClient.Country)

		for _, mail := range myClient.Inbox {
			fmt.Println(mail.RemoteEmail, ":", mail.Body)
		}
	})

	http.ListenAndServe(":3000", nil)
}