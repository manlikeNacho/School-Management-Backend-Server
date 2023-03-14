package models

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	CoursesTaken []Course `json:"courses_taken" gorm:"many2many:student_courses"`
}

type StudentResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
