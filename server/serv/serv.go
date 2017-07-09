package serv

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "Waxt/proto"
	"google.golang.org/grpc/reflection"
	env "Waxt/server/env"
)

const (
	port = ":1903"
)
var environment *env.Environment

type Server struct{}

func StartServer(en *env.Environment)  {

	environment = en

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterWaxtServer(s, &Server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
