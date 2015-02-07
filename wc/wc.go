package main

import (
	"bufio"
	"fmt"
	"os"
)

type WordCounter struct {
	LineCount int
}

func (wc *WordCounter) Line(line string) {
	wc.LineCount++
}

func main() {
	f, err := os.Open("the-prince.txt")
	if err != nil {
		panic(err)
	}
	
	var wc WordCounter
	
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wc.Line(scanner.Text())
	}
	
	fmt.Printf("Lines: %d\n", wc.LineCount)
}

