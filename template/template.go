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
{ {{range $key,$value := . }}
{ "{{$key}}" : "{{$value}}" },{{end}} }
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

type NameVal struct {
	Name string
	Val  string
}

func mapToKVPair(data map[string]string) []*NameVal {
	pairs := make([]*NameVal, 0)
	for k, v := range data {
		pair := new(NameVal)
		pair.Name = k
		pair.Val = v
		pairs = append(pairs, pair)
	}

	return pairs
}

//TODO - output to buffer, fix whitespace in buffer
func map3() {
	sampleTemplate :=
		`
[ {{range $index,$elem := . }}
	{{if $index}}
		,{ "{{$elem.Name}}" : "{{$elem.Val}}" }
	{{else}}
		{ "{{$elem.Name}}" : "{{$elem.Val}}" }
	{{end}}
{{end}}]
	`

	tmpl, err := template.New("test").Parse(sampleTemplate)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	kvpairs := make(map[string]string)
	kvpairs["foo"] = "Foo"
	kvpairs["bar"] = "Bar"

	data := mapToKVPair(kvpairs)

	tmpl.Execute(os.Stdout, data)
}

func main() {
	//map1()
	//map2()
	map3()
}
