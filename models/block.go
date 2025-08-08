package models

import "gorm.io/gorm"

type Block struct {
	gorm.Model
	Type    string `json:"type" gorm:"notnull"`
	Content string `json:"content" gorm:"notnull"`
	LevelID uint   `json:"levelID" gorm:"notnull"`
	Level   Level  `json:"level" gorm:"foreignKey:LevelID"`
}
