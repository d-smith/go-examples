package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type WordCounter struct {
	LineCount int
	WordCount int
}

func countWordsInLine(line string) int {
	wordCount := 0
	wordScanner := bufio.NewScanner(strings.NewReader(line))
	wordScanner.Split(bufio.ScanWords)
	for wordScanner.Scan() {
		wordCount++
	}

	return wordCount
}

func (wc *WordCounter) processLine(line string, countWords bool) {
	wc.LineCount++
	if countWords {
		wc.WordCount += countWordsInLine(line)
	}
}

func (wc *WordCounter) Count(f *os.File, countWords bool) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wc.processLine(scanner.Text(), countWords)
	}
}

func report(wc *WordCounter, lines bool, words bool) {
	if words && !lines {
		fmt.Printf("%d\n", wc.WordCount)
	} else if !words && lines {
		fmt.Printf("%d\n", wc.LineCount)
	} else {
		fmt.Printf("%d\t%d\n", wc.LineCount, wc.WordCount)
	}
}

func main() {
	linesPtr := flag.Bool("l", false, "count lines")
	wordsPtr := flag.Bool("w", false, "count words")
	flag.Parse()

	files := flag.Args()
	for i := range files {
		f, err := os.Open(files[i])
		if err != nil {
			panic(err)
		}

		var wc WordCounter
		wc.Count(f, *wordsPtr || (!*wordsPtr && !*linesPtr))
		f.Close()

		report(&wc, *wordsPtr, *linesPtr)
	}
}
