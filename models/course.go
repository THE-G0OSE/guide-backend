package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name      string `json:"name"`
	CreatorID uint   `json:"creator"`
	Levels    []uint `json:"levels"`
}

type GetCourseRequest struct {
	ID uint `json:"id"`
}

type DeleteCourseRequest struct {
	ID uint `json:"id"`
}

type CourseCreateRequest struct {
	Name string `json:"name"`
}

func RequestToCourse(in CourseCreateRequest, creator uint) Course {
	course := Course{
		Name:      in.Name,
		CreatorID: creator,
		Levels:    []uint{},
	}
	return course
}
