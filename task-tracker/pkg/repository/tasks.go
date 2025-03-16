package repository

import (
	"github.com/Domains18/subjective.git/config"
	"github.com/Domains18/subjective.git/internal/model"
	"gorm.io/gorm"
)


type task_response struct {
	Task *model.Task `json:"task"`
}

func GetAllTasks() ([]model.Task, error){
	var tasks []model.Task

	response := config.DB.Find(&tasks)

	if response.Error != nil {
		return nil, response.Error
	}

	return tasks, nil
}


func CreateTask(body model.Task)(task_response, error){
	if err := config.DB.Create(&body).Error; err != nil {
		return task_response{}, err
	}

	return task_response{
		Task: &body,
	}, nil
}



func FindTask(id string) (task_response, error) {
	var task model.Task 
	result := config.DB.Where("id = ?", id).First(&task)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return task_response{}, nil 
		}
		return task_response{}, result.Error 
	}

	return task_response{Task: &task}, nil
}
