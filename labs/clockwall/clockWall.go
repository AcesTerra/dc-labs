package main

import(
	//"fmt"
	"net"
	"log"
	"os"
	"io"
	//"time"
	"fmt"
)

func main() {
	var connections = os.Args
	fmt.Println(connections)
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
