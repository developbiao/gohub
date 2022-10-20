package user

import (
	"gohub/pkg/hash"
	"gorm.io/gorm"
)

// BeforeSave hook save hash password
func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}
