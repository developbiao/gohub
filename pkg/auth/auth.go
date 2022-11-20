package auth

import (
	"errors"
	"gohub/app/models/user"
)

// Attempt try to log in
func Attempt(email string, password string) (user.User, error) {
	userModel := user.GetByMulti(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("account does not exists")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("password Incorrect")
	}

	return userModel, nil
}

// LoginByPhone login by phone
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("phone not registered")
	}
	return userModel, nil
}
