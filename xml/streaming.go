package main

import (
	"encoding/xml"
	"strings"
	"fmt"
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
      </urn:AddNote>
   </soapenv:Body>
</soapenv:Envelope>
	`
	
	reader := strings.NewReader(xmlLiteral)
	decoder := xml.NewDecoder(reader)
	
	for {
		t,_ := decoder.Token()
		if t == nil {
			break
		}
		
		fmt.Println(t)
	}
	
}


