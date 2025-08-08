package models

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Name         string `json:"name"`
	Type         string `json:"type"`
	Position     []int  `json:"position"`
	Dependencies []uint `json:"dependencies"`
	Blocks       []uint `json:"blocks"`
}
