package models

import (
	"time"
)

type Statistics struct {
	UserID              uint      `gorm:"not null" json:"user_id"`
	CompletedTasksCount int       `gorm:"default:0" json:"completed_tasks_count"`
	TimerStartCount     int       `gorm:"default:0" json:"timer_start_count"`
	RecordDate          time.Time `gorm:"autoCreateTime" json:"record_date"`
}
