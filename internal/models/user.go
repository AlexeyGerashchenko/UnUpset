package models

import "time"

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Email          string    `gorm:"unique;not null" json:"email"`
	Password       string    `gorm:"not null" json:"password"`
	UserName       string    `gorm:"size:20" json:"user_name"`
	DateOfCreation time.Time `gorm:"autoCreateTime" json:"date_of_creation"`
}