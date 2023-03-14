package student

import (
	"context"
	"errors"

	"github.com/nurhidaylma/dts-intro-to-go/internal"
	"github.com/nurhidaylma/dts-intro-to-go/internal/domain"
	"github.com/nurhidaylma/dts-intro-to-go/internal/public"
)

func (s *StudentService) UpdateStudent(ctx context.Context, params *public.UpdateStudentRequest) (*public.StudentResponse, error) {
	var err error

	studentRepo, err := s.repository.FindByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	if studentRepo == nil {
		return nil, errors.New(internal.ErrInvalidRequest)
	}

	student := &domain.Student{}
	if params.Nama != "" {
		student.Nama = params.Nama
	}
	if params.Alamat != "" {
		student.Alamat = params.Alamat
	}
	if params.Pekerjaan != "" {
		student.Pekerjaan = params.Pekerjaan
	}
	if params.Motivasi != "" {
		student.Motivasi = params.Motivasi
	}

	updatedStudent, err := s.repository.Update(ctx, student.ToRepositoryModel())
	if err != nil {
		return nil, err
	}

	student.FromRepositoryModel(updatedStudent)

	return student.ToPublicModel(), nil
}
