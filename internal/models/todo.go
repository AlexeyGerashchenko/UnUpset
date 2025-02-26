package models

type ToDo struct {
	TaskID uint   `gorm:"primaryKey" json:"task_id"`
	UserID uint   `gorm:"not null" json:"user_id"`
	Text   string `gorm:"size:255;not null" json:"text"`
	Status bool   `gorm:"not null" json:"status"`
}