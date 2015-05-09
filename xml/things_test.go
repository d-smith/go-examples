package main

import (
	"testing"
)

//Run via go test -bench=BenchmarkXML1
func BenchmarkXML1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xml1(false)
	}
}

func BenchmarkStreamXML1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		streamParseXml1(false)
	}
}

func TestSonOfStreamParseXml1(t *testing.T) {
	result := sonOfStreamParseXml1(false)
	if result.Type != "Thing1" {
		t.Error("type should be Thing1")
	}

	if len(result.Things) != 18 {
		t.Error("expected 18 things")
	}

	if result.Things[0] != "xxx" {
		t.Error("First thing should be xxx")
	}

	if result.Things[1] != "yyy" {
		t.Error("Second thing should be yyy")
	}
}
