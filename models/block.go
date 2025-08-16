package models

import "gorm.io/gorm"

type Block struct {
	gorm.Model
	Type    string `json:"type"`
	Content string `json:"content"`
	LevelID uint   `json:"levelID"`
}
