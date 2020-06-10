package scheduler

import (
        "context"
        "log"
        "time"
	//"fmt"
	"strconv"
        //pb "github.com/CodersSquad/dc-labs/challenges/third-partial/proto"
        pb "github.com/AcesTerra/proto"
	"google.golang.org/grpc"
)

const (
      address = "localhost:50051"
//      defaultName = "world"
)

var idJob = 0

type Job struct {
        Address string
        RPCName string
}

func schedule(job Job, response chan string){
        // Set up a connection to the server.
        conn, err := grpc.Dial(job.Address, grpc.WithInsecure(), grpc.WithBlock())
        if err != nil {
                log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        c := pb.NewGreeterClient(conn)

        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        r, err := c.Test(ctx, &pb.HelloRequest{Name: job.RPCName})
        if err != nil {
                log.Fatalf("could not greet: %v", err)
        }
	idJob++
        log.Printf("Scheduler: RPC respose from %s : %s", job.Address, r.GetMessage())
	rpcResponse := r.GetMessage() + ";" + strconv.Itoa(idJob)
	//fmt.Println(rpcResponse)
	response <- rpcResponse
	//return rpcResponse
}

func Start(jobs chan Job, response chan string) error {
	for{
                job := <-jobs
                schedule(job, response)
        }
        return nil
}

