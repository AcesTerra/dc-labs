package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

func main() {
	fmt.Println("Pacman Log Analyzer")

	if len(os.Args) < 2 {
		fmt.Println("You must send at least one pacman log file to analize")
		fmt.Println("usage: ./pacman_log_analizer <logfile>")
		os.Exit(1)
	}

	// Your fun starts here.
	fmt.Println(os.Args[1])
	file, err := os.Open(os.Args[1])
    	if err != nil {
        	log.Fatal(err)
    	}
    	defer file.Close()

	var textLines []string

  	scanner := bufio.NewScanner(file)
    	for scanner.Scan() {
        	//fmt.Println(scanner.Text())
		textLines = append(textLines, scanner.Text())
    	}

    	if err := scanner.Err(); err != nil {
        	log.Fatal(err)
    	}

	for _, v := range textLines{
		fmt.Println(v)
	}
	
	testSplitedStr := strings.Split(textLines[0], " ")
	for _, v := range testSplitedStr {
		fmt.Println(v)
	}
}
