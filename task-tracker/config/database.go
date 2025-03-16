package config


import (
	"fmt"

	"github.com/Domains18/subjective.git/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var DB *gorm.DB

func Connect_DB() {
	var err error
	DB, err := gorm.Open(sqlite.Open("task_test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	DB.AutoMigrate(&model.Task{}, &model.Goals{}, &model.User{})
} 