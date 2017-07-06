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

func save(client pb.WaxtClient) {

	customer := &pb.Customer{Name: "yasin", Phone: "5416504615", Email: "vyasinw@gmail.com"}

	response, err := client.Save(context.Background(), customer)

	if err != nil {
		log.Fatalf("Could not saved customer: %v", err)
	}

	saved, _ := json.Marshal(response)

	log.Printf("Saved customer %v", string(saved))
}

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewWaxtClient(conn)

	save(client)
}
