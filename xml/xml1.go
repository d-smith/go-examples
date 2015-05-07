package main

import (
	"encoding/xml"
	"fmt"
)

func xml1(verbose bool) {
	var xmlDoc = `
	<foodoc>
		<type>Thing1</type>
		<things>
			<thing>xxx</thing>
			<thing>yyy</thing>
		</things>
	</foodoc>
	`

	type Result struct {
		XMLName xml.Name `xml:"foodoc"`
		Type    string   `xml:"type"`
		Things  []string `xml:"things>thing"`
	}

	v := new(Result)

	err := xml.Unmarshal([]byte(xmlDoc), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	if verbose {
		fmt.Printf("Type: %s\n", v.Type)
		fmt.Printf("Things: %v\n", v.Things)
	}

}
