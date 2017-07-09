package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "Waxt/proto"
	"encoding/json"
)

const (
	address = "localhost:1901"
)

func Save(client pb.WaxtClient) {

	customer := &pb.Customer{Id: int32(1), Name: "yasin", Phone: "5416504615", Email: "vyasinw@gmail.com"}

	response, err := client.Save(context.Background(), customer)

	if err != nil {
		log.Fatalf("Could not saved customer: %v", err)
	}

	saved, _ := json.Marshal(response)

	log.Printf("Save customer %v", string(saved))
}

func Get(client pb.WaxtClient) {
	response, err := client.Get(context.Background(), &pb.CustomerId{Id: int32(1)})

	if err != nil {
		log.Fatalf("Could not get customer: %v", err)
	}

	customer, _ := json.Marshal(response)

	log.Printf("Get customer %v", string(customer))
}

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewWaxtClient(conn)

	Get(client)
}
