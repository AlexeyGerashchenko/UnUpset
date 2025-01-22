package models

import (
	"time"
)

type ToDo struct {
	UserID uint   `gorm:"not null" json:"user_id"`
	TaskID uint   `gorm:"primaryKey" json:"task_id"`
	Text   string `gorm:"size:255;not null" json:"text"`
	Status bool   `gorm:"not null" json:"status"`
}
