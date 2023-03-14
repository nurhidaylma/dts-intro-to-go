package student

import (
	"context"

	"github.com/nurhidaylma/dts-intro-to-go/internal/domain"
)

func (s *StudentService) DeleteStudent(ctx context.Context, studentID int) error {
	studentData, err := s.GetStudent(ctx, studentID)
	if err != nil {
		return err
	}

	student := &domain.Student{}
	student.FromPublicModel(studentData)

	return s.repository.Delete(ctx, student.ToRepositoryModel())
}
