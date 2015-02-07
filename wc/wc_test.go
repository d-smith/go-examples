package main

import (
	"os"
	"testing"
)

func TestWordCounter(t *testing.T) {
	f, err := os.Open("the-prince.txt")
	if err != nil {
		t.Error("Error opening input file", err)
	} 
	
	var wc WordCounter
	wc.Count(f, true)
	f.Close()
	
	if(wc.WordCount != 52639) {
		t.Error("Expected 52639 words, got ", wc.WordCount)
	}
	
	if(wc.LineCount != 5063) {
		t.Error("Expected 5063 lines, got ", wc.LineCount)
	}
	
	
}