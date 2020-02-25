package main

import(
	//"fmt"
	"net"
	"log"
	"os"
	"io"
	//"time"
	"fmt"
	"strings"
)

func main() {
	var cmdLnArgs = os.Args
	var connections = cmdLnArgs[1:]
	//fmt.Println(connections)
	var ports []string
	for _, v := range connections{
		argsSplitted := strings.Split(v, "=")
		ports = append(ports, argsSplitted[1])
	}
	//fmt.Println(ports)
	c := make(chan int)
	for _, v := range ports{
		//conn, err := net.Dial("tcp", v)
		//if err != nil {
			//log.Fatal(err)
		//}
		//go mustCopy(os.Stdout, conn, c)
		go printHour(v, c)
	}
	info := <- c
	//defer conn.Close()
	fmt.Println(info)
}

func printHour(v string, c chan int) {
	conn, err := net.Dial("tcp", v)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	io.Copy(os.Stdout, conn)
	c <- 1
}
