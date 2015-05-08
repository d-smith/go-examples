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
