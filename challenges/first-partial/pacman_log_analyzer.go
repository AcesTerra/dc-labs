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

	var rawTextLines []string

  	scanner := bufio.NewScanner(file)
    	for scanner.Scan() {
        	//fmt.Println(scanner.Text())
		rawTextLines = append(rawTextLines, scanner.Text())
    	}

    	if err := scanner.Err(); err != nil {
        	log.Fatal(err)
    	}

	var filteredLines []string

	for _, v := range rawTextLines{
		//fmt.Println(v)
		testSplitedStr := strings.Split(v, " ")
		if testSplitedStr[3] == "installed" || testSplitedStr[3] == "upgraded" || testSplitedStr[3] == "removed"{
			filteredLines = append(filteredLines, v)
		}
	}
	
	//testSplitedStr := strings.Split(textLines[0], " ")
	for _, v := range filteredLines {
		fmt.Println(v)
	}

	//fmt.Println(testSplitedStr[3])
}
