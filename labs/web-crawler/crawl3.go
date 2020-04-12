// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"fmt"
	"log"
	"os"
	//"bufio"
	"gopl.io/ch5/links"
	"flag"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
var level = make(map[string]int)
var maxDepth int

func crawl(url string, depth int) ([]string, int) {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	depth++
	//fmt.Println(depth)
	//for _, v := range list{
	//	level[v] = depth
	//}
	return list, depth
}

//!-sema

//!+
func main() {
	worklist := make(chan []string)
	depChannel := make(chan int)
	var n int // number of pending sends to worklist
	dep := flag.Int("depth", 1, "an int")
	flag.Parse()
	maxDepth = *dep
	//fmt.Printf("maxDepth: %d\n", maxDepth)

	// Start with the command-line arguments.
	n++
	go func() {
		worklist <- os.Args[2:]
		depChannel <- 0
	}()
	level[os.Args[2]] = 0
	fmt.Println(os.Args[2])

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <- worklist
		depLink := <- depChannel
		//fmt.Println(depLink)
		for _, v := range list{
			fmt.Println(v)
			if _, ok := level[v]; ok {
				//continue
				//fmt.Println(level[v])
			} else{
				level[v] = depLink
			}
		}
		//bufio.NewReader(os.Stdin).ReadBytes('\n')
		for _, link := range list {
			//fmt.Printf("Link: %d\n", level[link])
			if level[link] < maxDepth {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						links, dep := crawl(link, level[link])
						worklist <- links
						depChannel <- dep
					}(link)
				}
			}
		}
	}
}

//!-
