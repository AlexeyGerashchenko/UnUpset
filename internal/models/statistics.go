package models

type Statistics struct {
	UserID              uint `gorm:"primaryKey" json:"user_id"`
	CompletedTasksCount int  `gorm:"default:0" json:"completed_tasks_count"`
	TimerStartCount     int  `gorm:"default:0" json:"timer_start_count"`
}