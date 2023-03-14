package student

import (
	"context"

	"github.com/nurhidaylma/dts-intro-to-go/internal/domain"
	"github.com/nurhidaylma/dts-intro-to-go/internal/public"
	"github.com/nurhidaylma/dts-intro-to-go/internal/repository"
)

// StudentServiceInterface represents the student service interface
type StudentServiceInterface interface {
	CreateStudent(ctx context.Context, params *domain.Student) (*public.StudentResponse, error)
	ListStudents(ctx context.Context, params *public.ListStudentRequest) ([]*public.StudentResponse, error)
	GetStudent(ctx context.Context, consultationID int) (*public.StudentResponse, error)
	UpdateStudent(ctx context.Context, params *public.UpdateStudentRequest) (*public.StudentResponse, error)
	DeleteStudent(ctx context.Context, consultationID int) error
}

// StudentService is the domain logic implementation of consultation service interface
type StudentService struct {
	repository repository.StudentRepository
}

// NewStudentService creates a new consultation domain Service
func NewStudentService(
	repository repository.StudentRepository,
) StudentServiceInterface {
	return &StudentService{
		repository: repository,
	}
}
