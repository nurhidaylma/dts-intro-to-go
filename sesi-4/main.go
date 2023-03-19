package main

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/nurhidaylma/dts-intro-to-go/config"
)

func main() {
	// Inject data to database
	injectSeed()

	arg := os.Args[1:]
	id, _ := strconv.Atoi(arg[0])

	student, err := findByID(id)
	if err == nil {
		fmt.Println(student)
	}
}

func injectSeed() {
	db := config.DB()

	_ = config.Seed(db)
}

// FindByID get student by student id
func findByID(studentID int) (*config.Student, error) {
	db := config.DB()

	student := config.Student{}
	if err := db.First(&student, `"id" = ? `, studentID).Error; err != nil {
		return nil, err
	}

	return &student, nil
}
