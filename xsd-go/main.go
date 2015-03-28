package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("XtracTypes.xsd")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := new(Schema)

	err = xml.Unmarshal(bs, &s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s.printSchemaStats()
}
