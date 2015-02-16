package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
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
