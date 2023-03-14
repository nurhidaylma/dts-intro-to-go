package postgres

import (
	"context"
	"time"

	"github.com/nurhidaylma/dts-intro-to-go/config"
	"github.com/nurhidaylma/dts-intro-to-go/internal/public"
	"github.com/nurhidaylma/dts-intro-to-go/internal/repository"
	"gorm.io/gorm"
)

// studentPostgres implements the student repository service
type studentPostgres struct {
	db *gorm.DB
	//FindAll collect all student data
}

func (s *studentPostgres) FindAll(ctx context.Context, params *public.ListStudentRequest) ([]repository.Student, error) {
	db := s.db

	students := []repository.Student{}
	args := []interface{}{}
	where := ``
	if len(params.IDs) > 0 {
		where += ` "id" IN ?`
		args = append(args, params.IDs)
	}
	if params.Search != "" {
		if where != "" {
			where += ` AND `
		}
		where += ` "nama" ILIKE ? OR "alamat" ILIKE ? OR "pekerjaan" ILIKE ?`
		args = append(args, "%"+params.Search+"%")
		args = append(args, "%"+params.Search+"%")
		args = append(args, "%"+params.Search+"%")
	}

	order := `"created_at" DESC`
	if err := db.Where(
		where,
		args...,
	).
		Order(order).
		Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}

// FindByID get student by student id
func (s *studentPostgres) FindByID(ctx context.Context, studentID int) (*repository.Student, error) {
	db := s.db

	student := repository.Student{}
	if err := db.First(&student, `"id" = ? `, studentID).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

// Insert inserts a student data
func (s *studentPostgres) Insert(ctx context.Context, student *repository.Student) (*repository.Student, error) {
	db := s.db

	now := time.Now().UTC()
	student.CreatedAt = now
	student.UpdatedAt = now

	if err := db.Create(student).Error; err != nil {
		return nil, err
	}

	return student, nil
}

// Update student data
func (s *studentPostgres) Update(ctx context.Context, updatedStudent *repository.Student) (*repository.Student, error) {
	db := s.db

	updatedStudent.UpdatedAt = time.Now().UTC()

	if err := db.Save(updatedStudent).Error; err != nil {
		return nil, err
	}

	return updatedStudent, nil
}

// Delete student data
func (s *studentPostgres) Delete(ctx context.Context, student *repository.Student) error {
	db := s.db

	if err := db.Delete(student).Error; err != nil {
		return err
	}

	return nil
}

// NewStudentPostgres creates new student repository service
func NewStudentPostgres() repository.StudentRepository {
	return &studentPostgres{
		db: config.DB(),
	}
}
