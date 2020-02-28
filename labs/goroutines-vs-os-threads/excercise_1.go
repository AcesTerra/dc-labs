package main

import(
	"fmt"
	"time"
	//"bufio"
	//"os"
)

func goRoutineLoop(returnToMain chan int, nextGoRoutine chan int, limit int){
	x := <- nextGoRoutine
	fmt.Println("Go routine number: ", x)
	if (x < limit){
		newGoRoutine := make(chan int)
		go goRoutineLoop(returnToMain, newGoRoutine, limit)
		x++
		newGoRoutine <- x
	} else{
		returnToMain <- x
	}
}

func main(){
	mainChannel := make(chan int)
	goRoutineChannel := make(chan int)
	start := 1
	var limit int
	_, err := fmt.Scanf("%d", &limit)
	if (err == nil){
		//fmt.Println(v)
		fmt.Println(time.Now().Format(time.RFC850))
		go goRoutineLoop(mainChannel, goRoutineChannel, limit)
		goRoutineChannel <- start
	}
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter a number: ")
	//text, _ := reader.ReadString('\n')
	//fmt.Println(text)
	<- mainChannel
	fmt.Println(time.Now().Format(time.RFC850))
	//fmt.Println(r)
}
