package model;

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`

	Goals []Goals `gorm:"foreignKey:UserId"`
	Tasks []Task `gorm:"foreignKey:UserId"`
}

type Task struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Description string `json:"description"`
	Status string `json:"status"`


	UserId uint `json:"user_id"`
	GoalId int `json:"goal_id"`
}

type Goals struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Description string `json:"description"`
	Status string `json:"status"`

	Tasks []Task `gorm:"foreignKey:GoalId"`
}