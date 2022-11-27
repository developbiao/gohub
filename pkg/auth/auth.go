package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
	"gohub/pkg/logger"
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

// CurrentUser get current  user
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("can not find user"))
		return user.User{}
	}
	// db is now a *DB value
	return userModel
}

// CurrentUID get current user id from gin.Context
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
