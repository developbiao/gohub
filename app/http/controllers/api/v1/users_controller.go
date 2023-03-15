package v1

import (
	"gohub/app/models/user"
	"gohub/pkg/auth"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// CurrentUser get current user information
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// Index get paginate data
func (ctrl *UsersController) Index(c *gin.Context) {
	data, paper := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": paper,
	})
}
