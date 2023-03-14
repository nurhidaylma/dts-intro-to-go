package repository

import (
	"context"
	"time"

	"github.com/nurhidaylma/dts-intro-to-go/internal/public"
)

type Student struct {
	ID        int       `json:"id" gorm:"primaryKey,not null,autoIncrement"`
	Nama      string    `json:"nama" gorm:"not null"`
	Alamat    string    `json:"alamat" gorm:"not null"`
	Pekerjaan string    `json:"pekerjaan" gorm:"not null"`
	Motivasi  string    `json:"motivasi" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null,autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null,autoUpdateTime"`
}

type StudentRepository interface {
	FindByID(ctx context.Context, studentID int) (*Student, error)
	FindAll(ctx context.Context, params *public.ListStudentRequest) ([]Student, error)
	Insert(ctx context.Context, student *Student) (*Student, error)
	Update(ctx context.Context, student *Student) (*Student, error)
	Delete(ctx context.Context, student *Student) error
}
