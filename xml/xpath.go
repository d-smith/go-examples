package main

import (
	"fmt"
	"gopkg.in/xmlpath.v1"
	"log"
	"reflect"
	"strings"
)

func xpathSampleFindOne(xmlDoc string, path string) {
	compiledPath := xmlpath.MustCompile(path)
	root, err := xmlpath.Parse(strings.NewReader(xmlDoc))
	if err != nil {
		log.Fatal(err)
	}

	if value, ok := compiledPath.String(root); ok {
		fmt.Println("type: ", reflect.TypeOf(value)) 
		fmt.Println("Found:", value)
	}
}

func xpathSampleFindMany(xmlDoc string, docpath string) {
	node, err := xmlpath.Parse(strings.NewReader(xmlDoc))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	path, err := xmlpath.Compile(docpath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var all []string
	iter := path.Iter(node)
	for iter.Next() {
		all = append(all, iter.Node().String())
	}

	fmt.Println(all)
}
