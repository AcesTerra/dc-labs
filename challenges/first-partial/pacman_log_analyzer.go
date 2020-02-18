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

//Function to check error in writing file
func checkError(err error) {
    if err != nil {
	panic(err)
    }
}

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

	//Creating the file to print report
	f, err := os.Create("packages_report.txt")
    	checkError(err)
    	defer f.Close()
	w := bufio.NewWriter(f)

	//Print headers
	title := "Pacman Log Analyzer"
	fmt.Println(title)
	drawLine(len(title))
	_, err = fmt.Fprintf(w, "Pacman Log Analizer\n")
	checkError(err)
	_, err = fmt.Fprintf(w, "-------------------\n")
	checkError(err)
	fmt.Printf("- Installed packages\t: %d\n", installedCntr)
	_, err = fmt.Fprintf(w, "- Installed packages\t: %d\n", installedCntr)
	checkError(err)
	fmt.Printf("- Removed packages\t: %d\n", removedCntr)
	_, err = fmt.Fprintf(w, "- Removeded packages\t: %d\n", removedCntr)
	checkError(err)
	fmt.Printf("- Upgraded packages\t: %d\n", upgradedCntr)
	_, err = fmt.Fprintf(w, "- Upgraded packages\t: %d\n", upgradedCntr)
	checkError(err)
	fmt.Printf("- Currently installed\t: %d\n", installedCntr - removedCntr)
	_, err = fmt.Fprintf(w, "- Currently installed\t: %d\n", installedCntr - removedCntr)
	checkError(err)
	subtitle := "List of packages"
	fmt.Println("\n" + subtitle)
	drawLine(len(subtitle))
	_, err = fmt.Fprintf(w, "\nList of packages\n")
	checkError(err)
	_, err = fmt.Fprintf(w, "----------------\n")
	checkError(err)

	//Print all packages and their info
	for _, i := range allPackages{
		fmt.Println("- Package name\t:", i.PackageName)
		_, err = fmt.Fprintf(w, "- Package name\t: %s\n", i.PackageName)
		checkError(err)
		fmt.Println("  - Install date\t:", i.InstallDate)
		_, err = fmt.Fprintf(w, "  - Install date\t: %s\n", i.InstallDate)
		checkError(err)
		fmt.Println("  - Last update date\t:", i.LastUpdate)
		_, err = fmt.Fprintf(w, "  - Last update date\t: %s\n", i.LastUpdate)
		checkError(err)
		fmt.Println("  - How many updates\t:", i.Upgrades)
		_, err = fmt.Fprintf(w, "  - How many updates\t: %d\n", i.Upgrades)
		checkError(err)
		fmt.Println("  - Removal date\t:", i.RemovalDate)
		_, err = fmt.Fprintf(w, "  - Removal date\t: %s\n", i.RemovalDate)
		checkError(err)
	}
	err = f.Close()
}
