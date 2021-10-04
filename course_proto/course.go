package course_proto

import (
	"context"
)

type server struct {
	UnimplementedCourseServiceServer
}

func (s *server) getCourse(ctx context.Context, id *Id) (*Course, error) {

	return &Course{ID: "1", Title: "Algo", Teacher: "Thore", Students: 100, Ects: 7.5, Raing: 9.1}, nil
}
