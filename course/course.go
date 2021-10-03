package course_proto

import (
	context "context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, course *Course) (*Course, error) {
	log.Printf("Received message from client: %s", course.Title)
	return &Course{Title: "Hello from server"}, nil
}
