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
	limit := 100
	go goRoutineLoop(mainChannel, goRoutineChannel, limit)
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter a number: ")
	//text, _ := reader.ReadString('\n')
	//fmt.Println(text)
	var v int
	_, err := fmt.Scanf("%d", &v)
	if (err == nil){
		//fmt.Println(v)
		fmt.Println(time.Now().Format(time.RFC850))
		goRoutineChannel <- v
	}
	r := <- mainChannel
	fmt.Println(time.Now().Format(time.RFC850))
	fmt.Println(r)
}
