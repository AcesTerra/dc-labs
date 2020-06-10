package scheduler

import (
        "context"
        "log"
        "time"
        "strconv"
        pb "github.com/AcesTerra/dc-labs/challenges/final/proto"
        "google.golang.org/grpc"
)

const (
      address = "localhost:50051"
//      defaultName = "world"
)

// Job ID counter
var idJob = 0

// Job structure
type Job struct {
        Address string
        RPCName string
}

// Ejecutes the job
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
        response <- rpcResponse
}

// Waits for a job to be executed
func Start(jobs chan Job, response chan string) error {
        for{
                job := <-jobs
                schedule(job, response)
        }
        return nil
}
