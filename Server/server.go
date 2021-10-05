package main

import (
	"context"
	"log"
	"net"

	pb "course_proto" // <- Insert into GOPATH

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

type server struct {
	pb.UnimplementedCourseServiceServer
}

var courses = []pb.Course{
	{ID: "1", Title: "Algo", Teacher: "Thore", Students: 100, Ects: 7.5, Raing: 9.1},
	{ID: "2", Title: "Code", Teacher: "Teach", Students: 200, Ects: 7.5, Raing: 8.9},
	{ID: "3", Title: "Design", Teacher: "Anders", Students: 50, Ects: 15, Raing: 8},
}

func (s *server) GetCourse(ctx context.Context, id *pb.Id) (*pb.Course, error) {
	log.Printf("Server received get command, ID: %s", id.GetID())
	for _, a := range courses {
		if a.GetID() == id.GetID() {
			log.Printf("Returning course: %s", &a)
			return &a, nil
		}
	}
	log.Printf("Course not found")
	return nil, nil
}
func (s *server) DeleteCourse(ctx context.Context, id *pb.Id) (*pb.Course, error) {
	log.Printf("Server received delete command, ID: %s", id.GetID())
	var index = 0
	for _, a := range courses {
		if a.GetID() == id.GetID() {
			log.Printf("Deleting course: %s", &a)
			courses = append(courses[:index], courses[index+1:]...)
			return &a, nil
		}
		index++
	}
	log.Printf("Course not found")
	return nil, nil
}
func main() {
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
