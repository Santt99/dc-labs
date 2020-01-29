package main

import (
	"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	var splitedString = strings.Split(string(s)," ")
	var words map[string]int
	words = map[string]int{} 
	for _, current := range splitedString {
		words[current]++
		fmt.Printf(string(words[current]))
	}
	return words
}

func main() {
	wc.Test(WordCount)
}


