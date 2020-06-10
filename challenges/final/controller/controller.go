package controller

import (
        "fmt"
        "os"
        "time"
        "strings"
        "go.nanomsg.org/mangos"
        "go.nanomsg.org/mangos/protocol/surveyor"
        _ "go.nanomsg.org/mangos/transport/all"
)

// Workload structure
type workload struct{
        name string
        jobId int
}

// Worker structure
type worker struct{
        name string
        tags string
        status string
        usage string
}

// Sockets for survey
var controllerAddress = "tcp://localhost:40899"
var sock mangos.Socket
var err error

// Slice to save registered workers
var allWorkers []worker

// Slice to store all workloads
var allWorkloads []workload

func die(format string, v ...interface{}) {
        fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
        os.Exit(1)
}

// Return date
func date() string {
        return time.Now().Format(time.ANSIC)
}

//Function to registry workload and check job id
func GetWorkloadJobId(workloadName string) int{
        isRegistered := false
        for i,v := range allWorkloads{
                if v.name == workloadName{
                        isRegistered = true
                        allWorkloads[i].jobId++
                        return allWorkloads[i].jobId
                }
        }
        if !isRegistered{
                registerWorkload := workload{name: workloadName, jobId: 1}
                allWorkloads = append(allWorkloads, registerWorkload)
                return 1
        }
        return 0
}

// Check worker info in memory
func WorkerStatus(workerName string) (string, string, string){
        for _,v := range allWorkers{
                if v.name == workerName{
                        return v.tags, v.status, v.usage
                }
        }
        return "","",""
}

// Survey architecture is used to check workers
func Start() {
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
}
