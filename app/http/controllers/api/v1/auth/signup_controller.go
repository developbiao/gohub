package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"
)

// SignupController  Signup controller
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist check phone is registered
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// Int Request object
	request := requests.SignupPhoneExistRequest{}

	// Validation check phone exists
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	// Check database exists phone number
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// Init request object
	request := requests.SignupEmailExistsRequest{}

	// Validation email exist
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	// Check data email exists
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	// Validation form
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// Create record
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()
	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后再试~")
	}
}
