package main

import (
	"log"
	"net"
	"os/exec"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func call_python(args []string) (string, error) {
	cmd := exec.Command("python3", "script.py", args[0])
	stdout, err := cmd.Output()
	if err != nil {
		println(err)
		return "", err
	}
	return string(stdout), nil
}

// --------------

type server struct {
	UnimplementedCipherServiceServer
}

func (s *server) Encode(ctx context.Context, message *CipherRequest) (*CipherResponse, error) {
	log.Printf("Recieved command: %s", message.GetName())
	return &CipherResponse{Body: "Hello from the server"}, nil
}

func main() {
	log.Println("Listening gRPC on 9000")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	grpcServer := grpc.NewServer()
	s := server{}
	RegisterCipherServiceServer(grpcServer, &s)
	log.Fatalln(grpcServer.Serve(lis))
}
