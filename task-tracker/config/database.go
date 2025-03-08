package config

import (
	"github.com/Domains18/subjective.git/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func Connect_DB() {
	db, err := gorm.Open(sqlite.Open("task_test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Task{}, &model.Goals{}, &model.User{})
}