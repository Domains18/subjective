package repository

import (
	"github.com/Domains18/subjective.git/config"
	"github.com/Domains18/subjective.git/internal/model"
	"github.com/Domains18/subjective.git/pkg/utils"
	"gorm.io/gorm"
)


type create_account_response struct {
	User *model.User `json:"user"`
	Error error `json:"error"`
	Message string `json:"message"`
	Token string `json:"token"`
}


func CreateAccount(body model.User) (create_account_response, error){
	exists, err := check_if_email_exists(body.Email)

	if err != nil{
		return create_account_response{
			User: nil,
			Error: err,
			Message: "failed to query user",
		}, err
	}

	if exists {
		return create_account_response{
			User: nil,
			Error: nil,
			Message: "user with the email already exists",
		}, err
	}

	hashed_password, err := utils.Hash_password(body.Password)

	if err != nil {
		return create_account_response{
			User: nil,
			Error: err,
			Message: "failed to hash password",
		}, err
	}

	user := model.User{
		Name: body.Name,
		Email: body.Email,
		Password: hashed_password,
	}
	

	if err := config.DB.Create(&user).Error; err != nil {
		return create_account_response{
			User: nil,
			Error: err,
			Message: "failed to create user",
		}, err
	}
	
	token, err := utils.Generate_JWT(user.ID, user.Email)
	if err != nil {
		return create_account_response{}, err
	}

	return create_account_response{
		User: &body,
		Error: nil,
		Message: "user created succesfully",
		Token: token,
	}, nil
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


func GetUsers() ([]model.User, error) {
	var users []model.User
	result := config.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}