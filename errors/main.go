package main

import (
	"fmt"
	"strings"
)

type MyError struct {
	Reasons []string
}

func NewMyError() *MyError {
	me := new(MyError)
	me.Reasons = make([]string, 2)
	return me
}

func (me *MyError) addReason(reason string) {
	me.Reasons = append(me.Reasons, reason)
}

func (me *MyError) Error() string {
	return fmt.Sprintf("%s", strings.Join(me.Reasons, " "))
}

func broken() (int, error) {
	me := NewMyError()
	me.addReason("firstly")
	me.addReason("secondly")
	me.addReason("thirdly")
	me.addReason("finally")
	return -1, me	
}

func main() {
	_, err := broken()
	if err != nil {
		panic(err)
	}	
}

