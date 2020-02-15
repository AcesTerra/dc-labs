package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	//"reflect"
)

func main() {
	fmt.Println("Pacman Log Analyzer")

	if len(os.Args) < 2 {
		fmt.Println("You must send at least one pacman log file to analize")
		fmt.Println("usage: ./pacman_log_analizer <logfile>")
		os.Exit(1)
	}

	// Open file
	//fmt.Println(os.Args[1])
	file, err := os.Open(os.Args[1])
    	if err != nil {
        	log.Fatal(err)
    	}
    	defer file.Close()

	//Variable that stores all lines of text
	var rawTextLines []string

	//Scan text lines
  	scanner := bufio.NewScanner(file)
    	for scanner.Scan() {
        	//fmt.Println(scanner.Text())
		rawTextLines = append(rawTextLines, scanner.Text())
    	}

    	if err := scanner.Err(); err != nil {
        	log.Fatal(err)
    	}

	//
	var filteredLines [][]string

	//Filter lines
	for _, v := range rawTextLines{
		//fmt.Println(v)
		splitedStr := strings.Split(v, " ")
		if splitedStr[3] == "installed" || splitedStr[3] == "upgraded" || splitedStr[3] == "removed"{
			filteredLines = append(filteredLines, splitedStr)
		}
	}
	
	//testSplitedStr := strings.Split(textLines[0], " ")
	//for _, v := range filteredLines {
		//fmt.Println(v)
	//}

	mapPackages := make(map[string][][]string)
	//x["key"] = append(x["key"], "value")
	//fmt.Println(filteredLines[0])
	//mapPackages[string(filteredLines[0][4])] = append(mapPackages[string(filteredLines[0][3])], filteredLines[0])
	fmt.Println(mapPackages)
	for _, v := range filteredLines{
		mapPackages[string(v[4])] = append(mapPackages[string(v[4])], v)
		//fmt.Println(v)
}
	//splitedStr := strings.Split(filteredLines)
	fmt.Println(mapPackages)
	//fmt.Println(reflect.TypeOf(mapPackages))
}
