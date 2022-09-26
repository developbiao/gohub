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
