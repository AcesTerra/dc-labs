package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	//"reflect"
)

type packageInfo struct{
	PackageName string
	InstallDate string
	LastUpdate string
	Upgrades int
	RemovalDate string
}

var installedCntr = 0
var upgradedCntr = 0
var removedCntr = 0

var allPackages []packageInfo
//var actualPackage packageInfo

func checkInfo(key string, logLine [][]string) {
	//actualPackage.InstallDate = "-"
	//actualPackage.LastUpdate = "-"
	//actualPackage.Upgrades = "-"
	//actualPackage.RemovalDate = "-"
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
			//upgradedCntr--
			removedCntr--
			installDate = "-"
			lastUpdate = "-"
			removalDate = "-"
			upgrades = 0
			hasBeenRemoved = false
			if i[3] == "installed"{
				installedCntr++
				installDate = i[0][1:] + " " + i[1][:len(i[1])-1]
				//installDate = i[0]
			}
			if i[3] == "upgraded"{
				//upgradedCntr++
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
				//installDate = i[0]
				//fmt.Println(reflect.TypeOf(i[0]))
			}
			if i[3] == "upgraded"{
				//upgradedCntr++
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

	//Filter lines
	for _, v := range rawTextLines{
		//fmt.Println(v)
		splitedStr := strings.Split(v, " ")
		if splitedStr[3] == "installed"{
			filteredLines = append(filteredLines, splitedStr)
			//installedCntr++
		}
		if splitedStr[3] == "upgraded"{
			filteredLines = append(filteredLines, splitedStr)
			//upgradedCntr++
		}
		if splitedStr[3] == "removed"{
			filteredLines = append(filteredLines, splitedStr)
			//removedCntr++
		}
	}

	//Map that contains the packages. Key is package name and stores a slice of slices of information
	mapPackages := make(map[string][][]string)
	//x["key"] = append(x["key"], "value")
	//fmt.Println(filteredLines[0])
	//mapPackages[string(filteredLines[0][4])] = append(mapPackages[string(filteredLines[0][3])], filteredLines[0])
	//fmt.Println(mapPackages)
	//Adding packages to mapPackages
	for _, v := range filteredLines{
		mapPackages[string(v[4])] = append(mapPackages[string(v[4])], v)
		//fmt.Println(v)
}
	//splitedStr := strings.Split(filteredLines)
	//fmt.Println(mapPackages["linux-firmware"][0][0][1:] + " " + mapPackages["linux-firmware"][0][1][:len(mapPackages["linux-firmware"][0][1])-1]) //Show date and time
	//fmt.Println(mapPackages["linux-firmware"][3][7][:len(mapPackages["linux-firmware"][3][7])-1]) //Show upgraded version
	//fmt.Println(mapPackages["python2"][0][5][1:len(mapPackages["python2"][0][5])-1]) //Show installed version
	//fmt.Println(reflect.TypeOf(mapPackages["ethtool"]))

	//k, v := mapPackages["linux-firmware"]
	//checkInfo(k, v)

	for i, v := range mapPackages{
		checkInfo(i, v)
	}

	title := "Pacman Log Analyzer"
	fmt.Println(title)
	drawLine(len(title))

	//Printing counters
	fmt.Printf("- Installed packages\t: %d\n", installedCntr)
	fmt.Printf("- Removeded packages\t: %d\n", removedCntr)
	fmt.Printf("- Upgraded packages\t: %d\n", upgradedCntr)
	fmt.Printf("- Currently installed\t: %d\n", installedCntr - removedCntr)

	subtitle := "List of packages"
	fmt.Println("\n" + subtitle)
	drawLine(len(subtitle))

	for _, i := range allPackages{
		fmt.Println("- Package name\t:", i.PackageName)
		fmt.Println("  - Install date\t:", i.InstallDate)
		fmt.Println("  - Last update date\t:", i.LastUpdate)
		fmt.Println("  - How many updates\t:", i.Upgrades)
		fmt.Println("  - Removal date\t:", i.RemovalDate)
	}

	/*for i, _ := range mapPackages{
		//fmt.Println("Key: ", i, "Value: ", v)
		fmt.Println("- Package Name\t: ", i)
		fmt.Println("   - Install date\t: ", mapPackages[i][0][0][1:] + " " + mapPackages[i][0][1][:len(mapPackages[i][0][1])-1]) //Show date and time)
		if len(mapPackages[i]) == 1{
			fmt.Println("   - Last update date\t: -")
			fmt.Println("   - How many updates\t: -")
			fmt.Println("   - Removal date\t: -")
		}
		//if len(mapPackages[i]) > 1{
			//fmt.Println("   - Last update date\t: ", mapPackages[i][len(mapPackages[i])-1][0][1:] + " " + mapPackages[i][len(mapPackages[i])-1][1][:len(mapPackages[i][len(mapPackages[i])-1][1])-1]) //Show date and time))
		//}
		if len(mapPackages[i]) > 1 && mapPackages[i][len(mapPackages[i])-1][3] == "removed" && mapPackages[i][len(mapPackages[i])-2][3] == "upgraded"{
			fmt.Println("   - Last update date\t: ", mapPackages[i][len(mapPackages[i])-2][0][1:] + " " + mapPackages[i][len(mapPackages[i])-2][1][:len(mapPackages[i][len(mapPackages[i])-2][1])-1])
			fmt.Printf("   - How many updates\t:  %d\n", len(mapPackages[i]) - 2)
			fmt.Println("   - Removal date\t: ", mapPackages[i][len(mapPackages[i])-1][0][1:] + " " + mapPackages[i][len(mapPackages[i])-2][1][:len(mapPackages[i][len(mapPackages[i])-2][1])-1])
		}
		if len(mapPackages[i]) > 1 && mapPackages[i][len(mapPackages[i])-1][3] == "removed" && mapPackages[i][len(mapPackages[i])-2][3] == "installed"{
			fmt.Println("   - Last update date\t: -")
			fmt.Println("   - How many updates\t: -")
			fmt.Println("   - Removal date\t: ", mapPackages[i][len(mapPackages[i])-1][0][1:] + " " + mapPackages[i][len(mapPackages[i])-2][1][:len(mapPackages[i][len(mapPackages[i])-2][1])-1])
		}
		if len(mapPackages[i]) > 1 && mapPackages[i][len(mapPackages[i])-1][3] == "upgraded"{
			fmt.Println("   - Last update date\t: ", mapPackages[i][len(mapPackages[i])-1][0][1:] + " " + mapPackages[i][len(mapPackages[i])-1][1][:len(mapPackages[i][len(mapPackages[i])-1][1])-1])
			fmt.Printf("   - How many updates\t:  %d\n", len(mapPackages[i]) - 1)
			fmt.Println("   - Removal date\t: -")
		}
		//fmt.Println(mapPackages["ethtool"][len(mapPackages["ethtool"])-1][3])
	}*/
	//fmt.Println(mapPackages["ethtool"][len(mapPackages["ethtool"])-1][3]) //Show action: installed, upgraded or removed
}
