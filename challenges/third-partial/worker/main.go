package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
	pb "github.com/AcesTerra/dc-labs/challenges/third-partial/proto"
	"go.nanomsg.org/mangos"
	"google.golang.org/grpc"
	"go.nanomsg.org/mangos/protocol/respondent"
	_ "go.nanomsg.org/mangos/transport/all"
)

var (
	defaultRPCPort = 50051
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// Worker info
var (
	controllerAddress = ""
	workerName = ""
	tags = ""
	status = "Running"
	usage = "50%"
)

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func date() string {
        return time.Now().Format(time.ANSIC)
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("RPC: Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// Test worker
func (s *server) Test(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("RPC Received. Job: %v\n",in.GetName())
	response := in.GetName() + ";" + status + ";" + "Done in worker " + workerName
	return &pb.HelloReply{Message: response}, nil
}

func init() {
	flag.StringVar(&controllerAddress, "controller", "tcp://localhost:40899", "Controller address")
	flag.StringVar(&workerName, "worker-name", "hard-worker", "Worker Name")
	flag.StringVar(&tags, "tags", "gpu,superCPU,largeMemory", "Comma-separated worker tags")
}

// joinCluster is meant to join the controller message-passing server
func joinCluster() {
	var sock mangos.Socket
	var err error
	var msg []byte

	if sock, err = respondent.NewSocket(); err != nil {
		die("can't get new respondent socket: %s", err.Error())
	}
	if err = sock.Dial(controllerAddress); err != nil {
		die("can't dial on respondent socket: %s", err.Error())
	}

	for {
		if msg, err = sock.Recv(); err != nil {
			die("Cannot recv: %s", err.Error())
		}
		fmt.Printf("Question: %s\n", string(msg))
		//d := date()
		fmt.Printf("Responfing my name\n")
		if err = sock.Send([]byte(workerName + ";" + tags + ";" + status + ";" + usage)); err != nil {
			die("Cannot send: %s", err.Error())
		}
	}
}

func getAvailablePort() int {
	port := defaultRPCPort
	for {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
		if err != nil {
			port = port + 1
			continue
		}
		ln.Close()
		break
	}
	return port
}

func main() {
	flag.Parse()

	// Subscribe to Controller
	go joinCluster()

	// Setup Worker RPC Server
	rpcPort := getAvailablePort()
	log.Printf("Starting RPC Service on localhost:%v", rpcPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", rpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
