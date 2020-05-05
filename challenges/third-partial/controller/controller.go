package controller

import (
	"fmt"
	//"log"
	"os"
	"time"

	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/protocol/pub"

	// register transports
	_ "go.nanomsg.org/mangos/transport/all"
)

var controllerAddress = "tcp://localhost:40899"
var sock mangos.Socket
var err error

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func date() string {
	return time.Now().Format(time.ANSIC)
}

func Start() {
	if sock, err = pub.NewSocket(); err != nil {
		die("can't get new pub socket: %s", err)
	}
	if err = sock.Listen(controllerAddress); err != nil {
		die("can't listen on pub socket: %s", err.Error())
	}
	//if err = sock.Dial(controllerAddress); err != nil {
	//	die("can't dial on sub socket: %s", err.Error())
	//}
	//workerConn, _ = sock.Recv()
	//fmt.Println(workerConn
	for {
		msg, msgErr := sock.Recv()
		if msgErr != nil{
			die("Cannot recv: %s", msgErr.Error())
		}
		fmt.Println(msg)
        }
	/*for {
		// Could also use sock.RecvMsg to get header
		d := date()
		log.Printf("Controller: Publishing Date %s\n", d)
		if err = sock.Send([]byte(d)); err != nil {
			die("Failed publishing: %s", err.Error())
		}
		time.Sleep(time.Second * 3)
	}*/
}
