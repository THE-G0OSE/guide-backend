package models

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Position     []int   `json:"position" gorm:"type:int[]"`
	Dependencies []Level `json:"dependencies" gorm:"many2many:level_level"`
	Blocks       []Block `json:"blocks"`
	CourseID     uint    `json:"courseID"`
	Course       Course  `json:"course" gorm:"foreignKey:CourseID"`
}
