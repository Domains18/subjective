package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Hash_password(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	if err != nil {
		return "", err
	}

		return string(hash), nil;
}


func Compare_password(hashed_password string, password string) (bool, error){
	if err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password));  err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return  false, err
	}
	return true, nil
}


func Generate_JWT(user_id uint, email string)(string, error){
	claims := jwt.MapClaims{
		"user_id": user_id,
		"email": email,
		"exp": time.Now().Add(time.Hour *24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString(os.Getenv("JWT_SECRET_KEY"))
}


func Validate_jwt(tokenString string)(*jwt.Token, error){
	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	return parsedToken, err
}