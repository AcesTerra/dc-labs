package main

import(
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string) //Creates a channel that can receive arrays of strings

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	//fmt.Println(len(worklist))
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
