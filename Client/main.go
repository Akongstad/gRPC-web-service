package main

import (
	"context"
	"log"
	"time"

	pb "course_proto" // <- Insert into GOPATH

	"google.golang.org/grpc"
)

const (
	address = ":9000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCourseServiceClient(conn)

	// Contact the server and print out its response.
	id := pb.Id{
		ID: "1",
	}

	clientMessage := "Client says: Give course with Id = " + id.ID
	log.Printf(clientMessage)
	time.Sleep(time.Second * 1)

	r, err := c.GetCourse(context.Background(), &id)
	if err != nil {
		log.Fatalf("could not find Course: %v", err)
	}
	log.Printf("Server says: %s", r)
	time.Sleep(time.Second * 1)

	id.ID = "2"
	clientMessage = "Client says: Give course with Id = " + id.ID
	log.Printf(clientMessage)
	time.Sleep(time.Second * 1)

	r2, err := c.GetCourse(context.Background(), &id)
	if err != nil {
		log.Fatalf("could not find Course: %v", err)
	}
	log.Printf("Server says: %s", r2)

	clientMessage = "Client says: Delete course with Id = " + id.ID
	log.Printf(clientMessage)
	id.ID = "1"
	time.Sleep(time.Second * 1)
	r3, err := c.DeleteCourse(context.Background(), &id)
	if err != nil {
		log.Fatalf("could not find Course: %v", err)
	}
	log.Printf("Server says: %s", r3)

	time.Sleep(time.Second * 1)
	clientMessage = "Client says: Give course with Id = " + id.ID
	log.Printf(clientMessage)
	r4, err := c.GetCourse(context.Background(), &id)
	if err != nil {
		log.Fatalf("could not find Course: %v", err)
	}
	log.Printf("Server says: %s", r4)
}
