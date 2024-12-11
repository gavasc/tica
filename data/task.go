package data

import (
	"gorm.io/gorm"
)

type Task struct {
  gorm.Model
  Code        string `gorm:"unique; not null;"`
  Description string
}

func (t Task) Create() error {
  db := connectDb()
	result := db.Create(&t)
	if result.Error != nil {
		return result.Error
	}
	db.Save(&t)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
