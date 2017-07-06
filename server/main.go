package main

import (
"log"
"net"

"golang.org/x/net/context"
"google.golang.org/grpc"
pb "Waxt/protoservice"
"google.golang.org/grpc/reflection"
	"fmt"
)

const (
	port = ":1903"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Save(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	fmt.Println("Customer",in)
	return in, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterWaxtServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

