package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name      *string `json:"name,omitempty" gorm:"notnull"`
	CreatorID uint    `json:"creator"`
	Levels    []Level `json:"levels,omitempty" gorm:"foreignKey:CourseID"`
}

type CourseCreateRequest struct {
	Name *string `json:"name"`
}

func RequestToCourse(in CourseCreateRequest, creator uint) Course {
	course := Course{
		Name:      in.Name,
		CreatorID: creator,
		Levels:    nil,
	}
	return course
}

type CoursePathcRequest struct {
	Name   *string `json:"name,omitempty"`
	Levels []Level `json:"levels,omitempty"`
}
