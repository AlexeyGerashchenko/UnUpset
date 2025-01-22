package models

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &ToDo{}, &Statistics{}, &Timer{})
	if err != nil {
		return err
	}
	return nil
}