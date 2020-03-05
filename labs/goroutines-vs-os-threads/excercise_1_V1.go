//V1: Send goroutines one after another with no end and without closing them.

package main

import(
	"fmt"
	"time"
	"os"
	"strconv"
)

//Function to check error in writing file
func checkError(err error) {
    if err != nil {
	panic(err)
    }
}

func goRoutineLoop(returnToMain chan int, nextGoRoutine chan int, f *os.File){
	x := <- nextGoRoutine
	fmt.Println("Go routine number:", x)
	goRoutineNumber := "Go routine number: " + strconv.Itoa(x) + "\n"
	_, errWrite := f.WriteString(goRoutineNumber)
	checkError(errWrite)
	x++
	go goRoutineLoop(returnToMain, nextGoRoutine, f)
	nextGoRoutine <- x
	<- returnToMain
	fmt.Println("Goroutine ended")
}

func main(){
	mainChannel := make(chan int)
	goRoutineChannel := make(chan int)
	start := 1
	f, err := os.Create("report_excercise_1_V1.txt")
    	checkError(err)
	startTime := time.Now().Format(time.RFC850) + "\n"
	fmt.Println(startTime)
	_, errWrite := f.WriteString(startTime)
	checkError(errWrite)
	go goRoutineLoop(mainChannel, goRoutineChannel, f)
	goRoutineChannel <- start
	defer f.Close()
	<- mainChannel
	endTime := time.Now().Format(time.RFC850) + "\n"
	fmt.Println(endTime)
	_, errWrite = f.WriteString(endTime)
	checkError(errWrite)
}
