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
	Blocks       []Block `json:"blocks" gorm:"foreignKey:LevelID"`
	CourseID     uint    `json:"courseID"`
}

type LevelPatchRequest struct {
	ID           *uint   `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Type         *string `json:"type,omitempty"`
	PositionX    *int    `json:"positionX,omitempty"`
	PositionY    *int    `json:"positionY,omitempty"`
	Dependencies []Level `json:"dependencies,omitempty"`
}
