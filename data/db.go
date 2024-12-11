package data

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data/data.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}

	db.AutoMigrate(&Task{}, &History{})

	return db
}
