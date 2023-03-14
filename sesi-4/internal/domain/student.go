package domain

import (
	"time"

	"github.com/nurhidaylma/dts-intro-to-go/internal/public"
	"github.com/nurhidaylma/dts-intro-to-go/internal/repository"
	"github.com/nurhidaylma/dts-intro-to-go/utils/helper"
)

type Student struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	Alamat    string    `json:"alamat"`
	Pekerjaan string    `json:"pekerjaan"`
	Motivasi  string    `json:"motivasi"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Student) FromPublicModel(studentPublic interface{}) {
	_ = helper.TransformObject(studentPublic, s)
}

func (s *Student) ToPublicModel() *public.StudentResponse {
	studentPublic := &public.StudentResponse{}
	_ = helper.TransformObject(s, studentPublic)
	return studentPublic
}

func (s *Student) FromRepositoryModel(consultationRepo interface{}) {
	_ = helper.TransformObject(consultationRepo, s)
}

func (s *Student) ToRepositoryModel() *repository.Student {
	consultationRepo := &repository.Student{}
	_ = helper.TransformObject(s, consultationRepo)
	return consultationRepo
}
