package auth

import (
	"errors"

	"github.com/Domains18/subjective.git/config"
	"github.com/Domains18/subjective.git/internal/model"
	"github.com/Domains18/subjective.git/pkg/utils"
)

type login_response struct {
	Message string      `json:"Message"`
	User    *model.User `json:"user"`
	Token   string      `json:"token"`
}

type login_body struct {
	email    string
	password string
}

func Login(body login_body) (login_response, error) {
	var user model.User

	if err := config.DB.Where("email = ?", body.email).First(&body).Error; err != nil {
		return login_response{}, errors.New("invalid email")
	}

	if match, err := utils.Compare_password(user.Password, body.password); err != nil {
		return login_response{}, err
	} else if !match {
		return login_response{Message: "wrong credentials", User: nil, Token: ""}, errors.New("invalid credentials")
	}

	token, err := utils.Generate_JWT(user.ID, user.Email)
	if err != nil {
		return login_response{}, err
	}
	return login_response{User: &user, Token: token, Message: "succesfully authenticated"}, nil
}
