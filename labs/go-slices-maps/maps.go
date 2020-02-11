package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var stringArray []string = strings.Fields(s)
	testMap := make(map[string]int)
	for _, word := range stringArray {
		count, ok := testMap[word]
		if ok {
			testMap[word] = count + 1
		} else {
			testMap[word] = 1
		}
	}
	return testMap
}

func main() {
	wc.Test(WordCount)
}
