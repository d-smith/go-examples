package main

import (
	"encoding/xml"
	"fmt"
)

func xml2() {
	var xmlDoc = `
	<foodoc>
		<type>Thing1</type>
		<things>
			<thing name='a'>xxx</thing>
			<thing name='b'>yyy</thing>
		</things>
	</foodoc>
	`
	type Thing struct {
		Name  string `xml:"name,attr"`
		Thing string `xml:",chardata"`
	}

	type Result struct {
		XMLName xml.Name `xml:"foodoc"`
		Type    string   `xml:"type"`
		Things  []Thing  `xml:"things>thing"`
	}

	v := new(Result)

	err := xml.Unmarshal([]byte(xmlDoc), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("Type: %s\n", v.Type)
	fmt.Printf("Things: %v\n", v.Things)

}
