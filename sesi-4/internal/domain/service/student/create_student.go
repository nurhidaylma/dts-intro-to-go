package student

import (
	"context"

	"github.com/nurhidaylma/dts-intro-to-go/internal/domain"
	"github.com/nurhidaylma/dts-intro-to-go/internal/public"
)

// CreateStudent create new student
func (s *StudentService) CreateStudent(ctx context.Context, params *domain.Student) (*public.StudentResponse, error) {

	var err error
	studentRepo := params.ToRepositoryModel()
	studentRepo, err = s.repository.Insert(ctx, studentRepo)
	if err != nil {
		return nil, err
	}

	params.FromRepositoryModel(studentRepo)
	studentResponse := params.ToPublicModel()

	return studentResponse, nil
}
