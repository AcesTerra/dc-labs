package main

import(
	"fmt"
	"time"
	//"bufio"
	"os"
	//"io"
	//"reflect"
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
	//fmt.Println("Go routine number:", x)
	goRoutineNumber := "Go routine number: " + strconv.Itoa(x) + "\n"
	_, errWrite := f.WriteString(goRoutineNumber)
	//err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	checkError(errWrite)
	//if (x < limit){
	x++
	//for{
		//newGoRoutine := make(chan int)
	go goRoutineLoop(returnToMain, nextGoRoutine, f)
		//x++
	nextGoRoutine <- x
	<- returnToMain
	fmt.Println("Goroutine ended")
	//}
}

func main(){
	mainChannel := make(chan int)
	goRoutineChannel := make(chan int)
	start := 1
	//var limit int
	//_, errInt := fmt.Scanf("%d", &limit)
	f, err := os.Create("report_excercise_1.txt")
    	checkError(err)
	//fmt.Println(reflect.TypeOf(f))
	//w := bufio.NewWriter(f)
	//if (errInt == nil){
		//fmt.Println(v)
	startTime := time.Now().Format(time.RFC850) + "\n"
	fmt.Println(startTime)
	_, errWrite := f.WriteString(startTime)
	checkError(errWrite)
	go goRoutineLoop(mainChannel, goRoutineChannel, f)
	goRoutineChannel <- start
	defer f.Close()
	//}
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter a number: ")
	//text, _ := reader.ReadString('\n')
	//fmt.Println(text)
	<- mainChannel
	//fmt.Println(time.Now().Format(time.RFC850))
	endTime := time.Now().Format(time.RFC850) + "\n"
	fmt.Println(endTime)
	_, errWrite = f.WriteString(endTime)
	checkError(errWrite)
	//fmt.Println(r)
}
