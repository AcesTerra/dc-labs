package main

import(
	"fmt"
	"os"
)

func main(){
	var env = os.Getenv("TZ")
	fmt.Println("TZ = ", env)
}
