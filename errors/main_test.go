package main

import (
	"testing"
	"strings"
)




func TestErrorReasons(t *testing.T) {
	_,err := broken()
	if(err == nil) {
		t.Error("Expected non-nil error message")
	} else {
		if !strings.Contains(err.Error(), "firstly") {
			t.Error("message should contain firstly")
		}
	}
}