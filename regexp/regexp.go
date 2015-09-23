package main

import (
	"regexp"
	"fmt"
)

var firstName = regexp.MustCompile(`^[a-zA-Z]+$`)
var lastName = regexp.MustCompile(`^([a-zA-Z'-]\s*)+$`)


func main() {
	fmt.Println(firstName.MatchString("name"))
	fmt.Println(firstName.MatchString("Name"))
	fmt.Println(firstName.MatchString("a name"))
	fmt.Println(firstName.MatchString("A name 123"))

	fmt.Println(lastName.MatchString("Smitty"))
	fmt.Println(lastName.MatchString("d'foobar"))
	fmt.Println(lastName.MatchString("funken-pitstain"))
	fmt.Println(lastName.MatchString("Sm1tty"))
	fmt.Println(lastName.MatchString("O Henry III"))
}