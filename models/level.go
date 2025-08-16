package models

import (
	"gorm.io/gorm"
)

type Level struct {
	gorm.Model
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	PositionX    int     `json:"positionX"`
	PositionY    int     `json:"positionY"`
	Dependencies []Level `json:"dependencies" gorm:"many2many:level_level"`
	Blocks       []Block `json:"blocks"`
	CourseID     uint    `json:"courseID"`
}
