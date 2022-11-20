package user

import "gohub/pkg/database"

// IsEmailExist check email is registered
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist check phone is registered
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByPhone find user by phone number
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// GetByMulti find user by phone/email/name
func GetByMulti(loginID string) (userModel User) {
	database.DB.Where("phone = ?", loginID).
		Where("email = ?", loginID).
		Where("name = ?", loginID).
		First(&userModel)
	return
}
