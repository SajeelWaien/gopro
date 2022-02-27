package main

import (
	"context"
	"log"

	"github.com/sajeelwaien/gopro/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := proto.NewSelectAgentClient(conn)

	message := proto.AgentName{
		Name: "Chamber",
	}
	res, err := client.SelectAgent(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling SelectAgent: %s", err)
	}

	log.Printf("Response: %s", res.Message)

}
