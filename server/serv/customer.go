package serv

import (
	"fmt"
	pb "Waxt/proto"
	"golang.org/x/net/context"
)



func (s *Server) Save(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	fmt.Println("Customer", in)

	return in, nil
}

func (s *Server) Get(ctx context.Context, in *pb.CustomerId) (*pb.Customer, error) {
	fmt.Println("Customer Id : ", in)

	return &pb.Customer{Id:in.Id,Name:"test",Email:"test@gmail.com",Phone:"1234567980"}, nil
}