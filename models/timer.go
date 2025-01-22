package models

import (
	"time"
)

type Timer struct {
	UserID      uint      `gorm:"not null" json:"user_id"`
	TimerLength int       `gorm:"not null" json:"timer_length"`
	StartTimer  time.Time `gorm:"not null" json:"start_timer"`
	EndTimer    time.Time `gorm:"not null" json:"end_timer"`
}
