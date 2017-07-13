package serv

import (
	pb "Waxt/proto"
	"golang.org/x/net/context"
)

func (s *Server) Save(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {

	db := environment.Store.Db

	err := db.Save(in)

	if err != nil {
		return nil, err
	}

	return in, nil
}

func (s *Server) Get(ctx context.Context, in *pb.CustomerId) (*pb.Customer, error) {

	var customer pb.Customer

	db := environment.Store.Db

	err := db.One("Id", in.Id, &customer)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}
