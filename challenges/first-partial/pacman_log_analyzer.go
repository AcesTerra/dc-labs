package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	//"reflect"
)

func drawLine (s int){
	for i :=  0; i < s; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

func main() {
	title := "Pacman Log Analyzer"
	fmt.Println(title)
	drawLine(len(title))

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

	//Lines that has installed, upgraded or removed
	var filteredLines [][]string
	installedCntr := 0
	upgradedCntr := 0
	removedCntr := 0

	//Filter lines
	for _, v := range rawTextLines{
		//fmt.Println(v)
		splitedStr := strings.Split(v, " ")
		if splitedStr[3] == "installed"{
			filteredLines = append(filteredLines, splitedStr)
			installedCntr++
		}
		if splitedStr[3] == "upgraded"{
			filteredLines = append(filteredLines, splitedStr)
			upgradedCntr++
		}
		if splitedStr[3] == "removed"{
			filteredLines = append(filteredLines, splitedStr)
			removedCntr++
		}
	}

	//Printing counters
	fmt.Printf("- Installed packages\t: %d\n", installedCntr)
	fmt.Printf("- Removeded packages\t: %d\n", removedCntr)
	fmt.Printf("- Upgraded packages\t: %d\n", upgradedCntr)
	fmt.Printf("- Currently installed\t: %d\n", installedCntr - removedCntr)

	subtitle := "List of packages"
	fmt.Println("\n" + subtitle)
	drawLine(len(subtitle))

	mapPackages := make(map[string][][]string)
	//x["key"] = append(x["key"], "value")
	//fmt.Println(filteredLines[0])
	//mapPackages[string(filteredLines[0][4])] = append(mapPackages[string(filteredLines[0][3])], filteredLines[0])
	//fmt.Println(mapPackages)
	for _, v := range filteredLines{
		mapPackages[string(v[4])] = append(mapPackages[string(v[4])], v)
		//fmt.Println(v)
}
	//splitedStr := strings.Split(filteredLines)
	fmt.Println(mapPackages["linux-firmware"][0][0][1:] + " " + mapPackages["linux-firmware"][0][1][:len(mapPackages["linux-firmware"][0][1])-1]) //Show date and time
	//fmt.Println(mapPackages["linux-firmware"][3][7][:len(mapPackages["linux-firmware"][3][7])-1]) //Show upgraded version
	//fmt.Println(mapPackages["python2"][0][5][1:len(mapPackages["python2"][0][5])-1]) //Show installed version
	//fmt.Println(mapPackages)
}
