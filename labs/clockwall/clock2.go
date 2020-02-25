// Clock2 is a concurrent TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
	"os"
	"fmt"
	"flag"
)

func TimeIn(t time.Time, timeZone string) (time.Time, error) {
	loc, err := time.LoadLocation(timeZone)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func handleConn(c net.Conn, timeZone string) {
	defer c.Close()
	for {
		t, err := TimeIn(time.Now(), timeZone)
		if err == nil {
			//fmt.Println(t.Location(), t.Format("15:04:05\n"))
		} else {
			fmt.Println(timeZone, "")
		}
		var locationAndTime = timeZone + "\t" + t.Format("15:04:05\n")
		_, er := io.WriteString(c, locationAndTime)
		if er != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//var port = os.Args[1]
	//var port = flag.String("port", "8080", "Eg: 9090")
	var port string
	flag.StringVar(&port, "port", "8080", "Eg: 9090")
	flag.Parse()
	port = "localhost:" + port
	var env = os.Getenv("TZ")
	//fmt.Println(env)
	//fmt.Println(port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, env) // handle connections concurrently
	}
}
