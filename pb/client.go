
package main

import (
	"github.com/golang/protobuf/proto"
	"net/http"
	"fmt"
	"bytes"
)

func main() {
	myClient := Client{Id: 526, Name: "John Doe", Email: "johndoe@example.com", Country: "US"}
	clientInbox := make([]*Client_Mail, 0, 20)
	clientInbox = append(clientInbox,
		&Client_Mail{RemoteEmail: "jannetdoe@example.com", Body: "Hello. Greetings. Bye."},
		&Client_Mail{RemoteEmail: "WilburDoe@example.com", Body: "Bye, Greetings, hello."})

	myClient.Inbox = clientInbox

	data, err := proto.Marshal(&myClient)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = http.Post("http://localhost:3000", "", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return
	}
}


