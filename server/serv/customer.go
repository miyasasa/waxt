package serv

import (
	pb "Waxt/proto"
	"golang.org/x/net/context"
	"github.com/boltdb/bolt"
	"encoding/json"
	"log"
	"strconv"
)

func (s *Server) Save(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {

	db := environment.Store.Db

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(environment.CustomerBucket))

		content, err := json.Marshal(in)

		if err != nil {
			log.Printf("Save marshal error %v", err)
			return err
		}

		err = bucket.Put([]byte(strconv.Itoa(int(in.Id))), content)

		if err != nil {
			log.Printf("Customer save failed %v", err)
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return in, nil
}

func (s *Server) Get(ctx context.Context, in *pb.CustomerId) (*pb.Customer, error) {

	var customer *pb.Customer

	db := environment.Store.Db

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(environment.CustomerBucket))

		content := bucket.Get([]byte(strconv.Itoa(int(in.Id))))

		err := json.Unmarshal(content, &customer)

		if err != nil {
			log.Printf("Get User Unmarshal error %v", err)
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return customer, nil
}
