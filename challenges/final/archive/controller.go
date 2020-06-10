package controller

import (
        "fmt"
        //"log"
        "os"
        "time"
	"strings"
        //"go.nanomsg.org/mangos"
        //"go.nanomsg.org/mangos/protocol/pub"
	"go.nanomsg.org/mangos"
	//"go.nanomsg.org/mangos/protocol/respondent"
	"go.nanomsg.org/mangos/protocol/surveyor"
        // register transports
        _ "go.nanomsg.org/mangos/transport/all"
)

type worker struct{
	name string
	tags string
	status string
	usage string
}

var controllerAddress = "tcp://localhost:40899"
var sock mangos.Socket
var err error
var allWorkers []worker

func die(format string, v ...interface{}) {
        fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
        os.Exit(1)
}

func date() string {
        return time.Now().Format(time.ANSIC)
}

func WorkerStatus(workerName string) (string, string, string){
	//Check worker info in memory
	for _,v := range allWorkers{
		if v.name == workerName{
			return v.tags, v.status, v.usage
		}
	}
	return "","",""
	//return status, tags, usage
}

func workerTest(){
	//Call scheduler
}

func Start() {
        //if sock, err = pub.NewSocket(); err != nil {
        //        die("can't get new pub socket: %s", err)
        //}
        //if err = sock.Listen(controllerAddress); err != nil {
        //        die("can't listen on pub socket: %s", err.Error())
        //}
        //if err = sock.Dial(controllerAddress); err != nil {
	//	die("can't dial on sub socket: %s", err.Error())
        //}

	//var sock mangos.Socket
	//var err error
	var msg []byte
	if sock, err = surveyor.NewSocket(); err != nil {
		die("can't get new surveyor socket: %s", err)
	}
	if err = sock.Listen(controllerAddress); err != nil {
		die("can't listen on surveyor socket: %s", err.Error())
	}
	err = sock.SetOption(mangos.OptionSurveyTime, time.Second)
	if err != nil {
		die("SetOption(): %s", err.Error())
	}

	for {
		time.Sleep(time.Second * 3)
		fmt.Println("Checking connection with workers")
		if err = sock.Send([]byte("Are you up?")); err != nil {
			die("Failed sending survey: %s", err.Error())
		}
		for {
			if msg, err = sock.Recv(); err != nil {
				break
			}
			isRegistered := false
			splitedMsg := strings.Split(string(msg), ";")
			registerWorker := worker{name: splitedMsg[0], tags: splitedMsg[1], status: splitedMsg[2], usage: splitedMsg[3]}
			for _,v := range allWorkers{
				if v.name == splitedMsg[0]{
					isRegistered = true
				}
			}
			if !isRegistered{
				allWorkers = append(allWorkers, registerWorker)
			}
			fmt.Printf("%s is up and running\n", string(splitedMsg[0]))
		}
		fmt.Println("Done checking")
	}

	/*for {
		//var msg []byte
		msg, msgErr := sock.Recv()
		if msgErr != nil{
			die("Cannot recv: %s", msgErr.Error())
		}
		fmt.Println(msg)
	}*/

        //workerConn, msgErr := sock.Recv()
	//if msgErr != nil{
	//	die("Cannot recv: %s", msgErr.Error())
	//}
        //fmt.Println(workerConn)
        /*for {
		//var msg []byte
		msg, msgErr := sock.Recv()
                if msgErr != nil{
			die("Cannot recv: %s", msgErr.Error())
		}
		fmt.Println(msg)
        }*/
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

