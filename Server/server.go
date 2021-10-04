package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "course_proto"

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

type server struct {
	pb.UnimplementedCourseServiceServer
}

func (s *server) GetCourse(ctx context.Context, id *pb.Id) (*pb.Course, error) {
	if id.GetID() == "1" {
		return &pb.Course{ID: "1", Title: "Algo", Teacher: "Thore", Students: 100, Ects: 7.5, Raing: 9.1}, nil

	} else {
		return nil, nil
	}
}

func main() {
	fmt.Println("Hello World!")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen port: %v", err)
	}
	s := server{}
	grpcServer := grpc.NewServer()

	pb.RegisterCourseServiceServer(grpcServer, &s)

	log.Printf("server listening at %v", lis.Addr())
	//fmt.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed serve server: %v", err)
	}
}
