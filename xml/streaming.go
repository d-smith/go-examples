package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func streaming() {

	xmlLiteral := `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:schemas-xtrac-fmr-com:b2b">
   <soapenv:Header>
	<urn:Cookie>261765034988290296725527451071864505174</urn:Cookie>
   </soapenv:Header>
   <soapenv:Body>
      <urn:AddNote>
         <urn:WorkItemNo>W019039-27NOV01</urn:WorkItemNo>
         <urn:Name></urn:Name>
         <urn:ControlNo>123</urn:ControlNo>
         <urn:Memo>some notes</urn:Memo>
         <urn:Note>c e-flat f-sharp a</urn:Note>
         <urn:Foo a='1' b='2'/>
      </urn:AddNote>
   </soapenv:Body>
</soapenv:Envelope>
	`

	reader := strings.NewReader(xmlLiteral)
	decoder := xml.NewDecoder(reader)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch t.(type) {
		default:
			fmt.Println("Not sure what we're dealing with here")
		case xml.StartElement:
			fmt.Println("StartElement")
			fmt.Printf("\tElement: %s\n", t.(xml.StartElement).Name)
			attrs := t.(xml.StartElement).Attr
			for _,attr := range attrs {
				fmt.Printf("\t attr: %s\n", attr)
			}
		case xml.EndElement:
			fmt.Println("EndElement")
		case xml.CharData:
			fmt.Println("CharData")
		case xml.Comment:
			fmt.Println("Comment")
		case xml.ProcInst:
			fmt.Println("ProcInst")
		case xml.Directive:
			fmt.Println("Directive")
		}
	}

}
