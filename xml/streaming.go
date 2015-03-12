package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)


func streaming(xmlLiteral string) {

	

	reader := strings.NewReader(xmlLiteral)
	decoder := xml.NewDecoder(reader)
	depth := 1
	maxDepth := 1
	startElement := false
	endElement := false

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
			
			endElement = false
			if startElement == true {
				
				depth++
				fmt.Println("start element true, depth now ", depth)
				if depth > maxDepth {
					maxDepth++
				}
			} else {
				startElement = true
			}
			
			attrs := t.(xml.StartElement).Attr
			for _,attr := range attrs {
				fmt.Printf("\t attr: %s\n", attr)
			}
		case xml.EndElement:
			fmt.Println("EndElement")
			startElement = false
			if endElement == true {
				depth--
			} else {
				endElement = true
			}
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
	
	fmt.Println("max depth: ", maxDepth)

}
