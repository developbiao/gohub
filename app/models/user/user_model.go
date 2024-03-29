package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"gohub/pkg/hash"
)

type User struct {
	models.BaseModel

	City         string `json:"city,omitempty"`
	Introduction string `json:"introduction,omitempty""`
	Avatar       string `json:"avatar,omitempty"`

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// ComparePassword check password is match
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// Save user record return rows affected
func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
