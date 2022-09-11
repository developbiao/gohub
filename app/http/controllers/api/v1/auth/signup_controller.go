package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"net/http"
)

// SignupController  Signup controller
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist check phone is registered
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// Request object
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}
	request := PhoneExistRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		// Parsing faire return 422 status code
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// Print error
		fmt.Println(err.Error())
		return
	}

	// Check database exists phone number
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
