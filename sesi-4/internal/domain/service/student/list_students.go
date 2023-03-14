package student

import (
	"context"

	"github.com/nurhidaylma/dts-intro-to-go/internal/domain"
	"github.com/nurhidaylma/dts-intro-to-go/internal/public"
)

func (s *StudentService) ListStudents(ctx context.Context, params *public.ListStudentRequest) ([]*public.StudentResponse, error) {
	studentRepo, err := s.repository.FindAll(ctx, params)
	if err != nil {
		return nil, err
	}

	result := []*public.StudentResponse{}
	for _, student := range studentRepo {
		studentDomain := &domain.Student{}
		studentDomain.FromRepositoryModel(student)

		studentResponse := studentDomain.ToPublicModel()

		result = append(result, studentResponse)
	}

	return result, nil
}
