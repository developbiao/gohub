package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
)

type LoginController struct {
	v1.BaseAPIController
}

// LoginByPhone login user with phone number
func (lc *LoginController) LoginByPhone(c *gin.Context) {
	// 1. Validation form
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	// 2. Attempt to login by phone
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// Login failed
		response.Error(c, err, "Account does not exist")
	} else {
		// Login success
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
