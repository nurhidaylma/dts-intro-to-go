package config

import "gorm.io/gorm"

type Student struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Pekerjaan string `json:"pekerjaan"`
	Motivasi  string `json:"motivasi"`
}

func seedUp() []interface{} {
	result := []interface{}{
		&Student{
			ID:        1,
			Nama:      "Naufal Hanif",
			Alamat:    "Jakarta",
			Pekerjaan: "Backend Engineer",
			Motivasi:  "Ingin memperkaya skill",
		}, &Student{
			ID:        2,
			Nama:      "Ina Nur",
			Alamat:    "Jakarta",
			Pekerjaan: "Backend Engineer",
			Motivasi:  "Ingin memperkaya skill",
		}, &Student{
			ID:        3,
			Nama:      "Patar Martua",
			Alamat:    "Jakarta",
			Pekerjaan: "Backend Engineer",
			Motivasi:  "Ingin memperkaya skill",
		}, &Student{
			ID:        4,
			Nama:      "Steven Setiawan",
			Alamat:    "Jakarta",
			Pekerjaan: "Backend Engineer",
			Motivasi:  "Ingin memperkaya skill",
		}, &Student{
			ID:        5,
			Nama:      "Ridha Ansari",
			Alamat:    "Jakarta",
			Pekerjaan: "Backend Engineer",
			Motivasi:  "Ingin memperkaya skill",
		}, &Student{
			ID:        6,
			Nama:      "Ilma",
			Alamat:    "Jakarta",
			Pekerjaan: "Backend Engineer",
			Motivasi:  "Ingin memperkaya skill",
		},
	}
	return result
}

func Seed(db *gorm.DB) error {
	seeders := seedUp()
	for _, seed := range seeders {
		err := db.Create(seed).Error
		if err != nil {
			return err
		}
	}
	return nil
}
