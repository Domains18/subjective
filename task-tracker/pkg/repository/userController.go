package repository

import (
	"net/http"

	"github.com/Domains18/subjective.git/internal/model"
	"gorm.io/gorm"
	
)

var db *gorm.DB

func CreateAccount(body model.User) (*model.User, error){
	err, exists := check_if_email_exists(body.Email)

	if err != nil{
		
	}
}



func check_if_email_exists(email string) (error, bool){
	exists := db.Where("email = ?", email).First(&email)

	if exists.Error != nil{
		return nil, false
	}
	println(exists)
	return nil, true
}