package main

import (
	"gopkg.in/xmlpath.v1"
	"log"
	"fmt"
	"strings"
)



func xpathSamples(xmlDoc string, path string) {
	compiledPath := xmlpath.MustCompile(path)
	root, err := xmlpath.Parse(strings.NewReader(xmlDoc))
	if err != nil {
		log.Fatal(err)
	}
	
	
	
	if value, ok := compiledPath.String(root); ok {
		fmt.Println("Found:", value)
	}
}
