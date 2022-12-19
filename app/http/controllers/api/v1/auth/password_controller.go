package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"
)

// PasswordController user controller
type PasswordController struct {
	v1.BaseAPIController
}

// ResetByPhone reset password by user phone
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	// 1. Validation form
	request := requests.RestByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.RestByPhone); !ok {
		return
	}

	// 2. Update password
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}
