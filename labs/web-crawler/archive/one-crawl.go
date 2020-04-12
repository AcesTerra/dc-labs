package main

import(
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main(){
	//url := "http://www.gopl.io/"
	url := os.Args[1]
	testString := crawl(url)
	for _, v := range testString{
		fmt.Println(v)
	}
}
