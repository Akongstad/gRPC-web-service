package main

import (
	"fmt"
	"log"
	"net"

	//pb "github.com/Akongstad/gRPC-web-service"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

/* type server struct {
	pb.UnimplementedGreeterServer
} */

/* func (s *Server) getCourse(ctx context.Context, course *Course) (*Course, error) {
	log.Printf("Received message from client: %s", course.Title)
	return &Course{Title: "Hello from server"}, nil
} */
func main() {
	fmt.Println("Hello World!")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen port 9000: %v", err)
	}
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed serve server: %v", err)
	}
}
