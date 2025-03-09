package repository

import (
	"github.com/Domains18/subjective.git/config"
	"github.com/Domains18/subjective.git/internal/model"
	"gorm.io/gorm"
)


type create_account_response struct {
	User *model.User `json:"user"`
	Error error `json:"error"`
	Message string `json:"message"`
}


func CreateAccount(body model.User) create_account_response{
	exists, err := check_if_email_exists(body.Email)

	if err != nil{
		return create_account_response{
			User: nil,
			Error: err,
			Message: "failed to query user",
		}
	}

	if exists {
		return create_account_response{
			User: nil,
			Error: nil,
			Message: "user with the email already exists",
		}
	}

	if err := config.DB.Create(&body).Error; err != nil {
		return create_account_response{
			User: nil,
			Error: err,
			Message: "failed to create user",
		}
	} 

	return create_account_response{
		User: &body,
		Error: nil,
		Message: "user created succesfully",
	}
}



func check_if_email_exists(email string) (bool, error){
	result := config.DB.Where("email = ?", email).First(&email)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound{
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}