package student

import (
	"context"

	"github.com/nurhidaylma/dts-intro-to-go/internal/domain"
	"github.com/nurhidaylma/dts-intro-to-go/internal/public"
)

func (s *StudentService) GetStudent(ctx context.Context, studentID int) (*public.StudentResponse, error) {
	student := &domain.Student{}
	studentRepo, err := s.repository.FindByID(ctx, studentID)
	if err != nil {
		return nil, err
	}

	student.FromRepositoryModel(studentRepo)
	studentResponse := student.ToPublicModel()

	return studentResponse, nil
}
