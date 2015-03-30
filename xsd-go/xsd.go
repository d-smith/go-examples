package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Schema struct {
	XMLName            xml.Name      `xml:"http://www.w3.org/2001/XMLSchema schema"`
	XT                 string        `xml:"xmlns xt,attr"`
	TargetNamespace    string        `xml:"targetNamespace,attr"`
	ElementFormDefault string        `xml:"elementFormDefault,attr"`
	Version            string        `xml:"version,attr"`
	Elements           []Element     `xml:"http://www.w3.org/2001/XMLSchema element"`
	ComplexTypes       []ComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	SimpleTypes        []SimpleType  `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
}

func (s *Schema) printSchemaStats() {
	fmt.Printf("xmlns xt: %s\n", s.XT)
	fmt.Printf("Target namespace: %s\n", s.TargetNamespace)
	fmt.Printf("Elements: %d\n", len(s.Elements))
	fmt.Printf("Complex type: %d\n", len(s.ComplexTypes))
	fmt.Printf("Simple type: %d\n", len(s.SimpleTypes))
}

type Element struct {
	XMLName      xml.Name     `xml:"http://www.w3.org/2001/XMLSchema element"`
	Type         string       `xml:"type,attr"`
	Ref          string       `xml:"ref,attr"`
	Nillable     string       `xml:"nillable,attr"`
	MinOccurs    string       `xml:"minOccurs,attr"`
	MaxOccurs    string       `xml:"maxOccurs,attr"`
	Form         string       `xml:"form,attr"`
	Name         string       `xml:"name,attr"`
	ComplexTypes *ComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
}

type ComplexType struct {
	XMLName  xml.Name        `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	Name     string          `xml:"name,attr"`
	Abstract bool            `xml:"abstract,attr"`
	Sequence []Element       `xml:"sequence>element"`
}

type SimpleType struct {
	XMLName     xml.Name    `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
	Name        string      `xml:"name,attr"`
	Constraints Restriction `xml:"http://www.w3.org/2001/XMLSchema restriction"`
}

type Restriction struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema restriction"`
	Base    string   `xml:"base,attr"`
}

func IsQualifiedName(name string) bool {
	return (strings.Count(name, ":") == 1)
}

func UnqualifiedName(name string) string {
	if !IsQualifiedName(name) {
		return name
	} else {
		parts := strings.Split(name, ":")
		return parts[1]
	}
}

func HasXsdPrefix(name string) bool {
	if !IsQualifiedName(name) {
		return false
	}
	
	parts := strings.Split(name, ":")
	
	return parts[0] == "xsd"
}






