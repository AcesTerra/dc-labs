/*
	Author: Alfredo Emmanuel Garcia Falcon
	Class: Distributed Computing
	Proffesor: Obed Muñoz
	Assignment: First Partial
	Title: Pacman Log Analizer
	Github: https://github.com/AcesTerra/dc-labs.git
*/

package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

//Structure to store package information
type packageInfo struct{
	PackageName string
	InstallDate string
	LastUpdate string
	Upgrades int
	RemovalDate string
}

//Global counter variables
var installedCntr = 0
var upgradedCntr = 0
var removedCntr = 0

//Slice of structures that stores all packages and their info
var allPackages []packageInfo

func checkInfo(key string, logLine [][]string) {
	var actualPackage packageInfo
	var upgrades = 0
	var installDate = "-"
	var lastUpdate = "-"
	var removalDate = "-"
	var hasBeenUpgraded = false
	var hasBeenRemoved = false
	actualPackage.PackageName = key
	for _, i := range logLine{
		if hasBeenRemoved == true {
			installedCntr--
			removedCntr--
			installDate = "-"
			lastUpdate = "-"
			removalDate = "-"
			upgrades = 0
			hasBeenRemoved = false
			if i[3] == "installed"{
				installedCntr++
				installDate = i[0][1:] + " " + i[1][:len(i[1])-1]
			}
			if i[3] == "upgraded"{
				hasBeenUpgraded = true
				upgrades++
				lastUpdate = i[0][1:] + " " + i[0][:len(i[1])-1]
			}
			if i[3] == "removed"{
				hasBeenRemoved = true
				removedCntr++
				removalDate = i[0][1:] + " " + i[1][:len(i[1])-1]
			}
		} else {
			if i[3] == "installed"{
				installedCntr++
				installDate = i[0][1:] + " " + i[1][:len(i[1])-1]
			}
			if i[3] == "upgraded"{
				hasBeenUpgraded = true
				upgrades++
				lastUpdate = i[0][1:] + " " + i[1][:len(i[1])-1]
			}
			if i[3] == "removed"{
				hasBeenRemoved = true
				removedCntr++
				removalDate = i[0][1:] + " " + i[1][:len(i[1])-1]
			}
		}
	}
	if hasBeenUpgraded{
		upgradedCntr++
	}
	actualPackage.InstallDate = installDate
	actualPackage.LastUpdate = lastUpdate
	actualPackage.Upgrades = upgrades
	actualPackage.RemovalDate = removalDate
	allPackages = append(allPackages, actualPackage)
}

//Function that creates a line given a number
func drawLine (s int){
	for i :=  0; i < s; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("You must send at least one pacman log file to analize")
		fmt.Println("usage: ./pacman_log_analizer <logfile>")
		os.Exit(1)
	}

	// Open file
	file, err := os.Open(os.Args[1])
    	if err != nil {
        	log.Fatal(err)
    	}
    	defer file.Close()

	//Slice that stores all lines of text file
	var rawTextLines []string

	//Scan text lines
  	scanner := bufio.NewScanner(file)
    	for scanner.Scan() {
		rawTextLines = append(rawTextLines, scanner.Text())
    	}

    	if err := scanner.Err(); err != nil {
        	log.Fatal(err)
    	}

	//Slice that contains text lines after filter
	var filteredLines [][]string

	//Filter raw text lines
	for _, v := range rawTextLines{
		splitedStr := strings.Split(v, " ")
		if splitedStr[3] == "installed" || splitedStr[3] == "upgraded" || splitedStr[3] == "removed"{
			filteredLines = append(filteredLines, splitedStr)
		}
	}

	//Map that contains the packages. Key is package name and stores a slice of slices of information
	mapPackages := make(map[string][][]string)

	//Fill the map
	for _, v := range filteredLines{
		mapPackages[string(v[4])] = append(mapPackages[string(v[4])], v)
	}

	//Analize map
	for i, v := range mapPackages{
		checkInfo(i, v)
	}

	//Print headers
	title := "Pacman Log Analyzer"
	fmt.Println(title)
	drawLine(len(title))
	fmt.Printf("- Installed packages\t: %d\n", installedCntr)
	fmt.Printf("- Removeded packages\t: %d\n", removedCntr)
	fmt.Printf("- Upgraded packages\t: %d\n", upgradedCntr)
	fmt.Printf("- Currently installed\t: %d\n", installedCntr - removedCntr)
	subtitle := "List of packages"
	fmt.Println("\n" + subtitle)
	drawLine(len(subtitle))

	//Print all packages and their info
	for _, i := range allPackages{
		fmt.Println("- Package name\t:", i.PackageName)
		fmt.Println("  - Install date\t:", i.InstallDate)
		fmt.Println("  - Last update date\t:", i.LastUpdate)
		fmt.Println("  - How many updates\t:", i.Upgrades)
		fmt.Println("  - Removal date\t:", i.RemovalDate)
	}
}
