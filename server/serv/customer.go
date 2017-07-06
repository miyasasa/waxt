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
