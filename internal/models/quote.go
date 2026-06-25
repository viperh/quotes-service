package models

import "gorm.io/gorm"

type Quote struct {
	gorm.Model
	Text string `gorm:"uniqueIndex;not null" json:"text"`
}
