package main

import(
	"fmt"
	"time"
	"os"
	"strconv"
)

var pingPongCntr = 0

//Function to check error in writing file
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ping(pingChannel chan bool, pongChannel chan bool, f *os.File){
	for range pingChannel {
		pingPongCntr++
		strCntr := strconv.Itoa(pingPongCntr) + "\n"
		_, errWrite := f.WriteString(strCntr)
		checkError(errWrite)
		pongChannel <- true
	}
}

func pong(pingChannel chan bool, pongChannel chan bool, f *os.File){
	for range pongChannel {
		pingPongCntr++
		strCntr := strconv.Itoa(pingPongCntr) + "\n"
		_, errWrite := f.WriteString(strCntr)
		checkError(errWrite)
		pingChannel <- true
	}
}

func main(){
	pingChannel := make(chan bool)
	pongChannel := make(chan bool)
	//start := 1
	f, err := os.Create("report_excercise_2.txt")
    	checkError(err)
	startTime := time.Now().Format(time.RFC850) + "\n"
	fmt.Println(startTime)
	_, errWrite := f.WriteString(startTime)
	checkError(errWrite)
	go ping(pingChannel, pongChannel, f)
	go pong(pingChannel, pongChannel, f)
	pingChannel <- true
	//goRoutineChannel <- start
	time.Sleep(time.Second * 1)
	//<- mainChannel
	endTime := time.Now().Format(time.RFC850) + "\n"
	fmt.Println(endTime)
	_, errWrite = f.WriteString(endTime)
	checkError(errWrite)
	defer f.Close()
}
