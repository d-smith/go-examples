package main

import (
	"fmt"
	"html/template"
	"os"
)

func map1() {
	sampleTemplate := `{ "{{.foo}}" : "{{.bar}}" }`

	tmpl, err := template.New("test").Parse(sampleTemplate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	data := make(map[string]string)
	data["foo"] = "Foo"
	data["bar"] = "Bar"

	tmpl.Execute(os.Stdout, data)
}

func map2() { 
	//TODO - get the commas right via use of array and index=0 test
	sampleTemplate := 
	`
{ {{range $key,$value := . }}{ "{{$key}}" : "{{$value}}" },{{end}} }
	`
	tmpl, err := template.New("test").Parse(sampleTemplate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	data := make(map[string]string)
	data["foo"] = "Foo"
	data["bar"] = "Bar"
	
	tmpl.Execute(os.Stdout, data)
 }

func main() {
	map1()
	map2()
}
