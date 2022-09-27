package response

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/logger"
	"gorm.io/gorm"
	"net/http"
)

// JSON response 200 and json data
func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// Success response 200 success and default data
func Success(c *gin.Context, data ...interface{}) {
	JSON(c, gin.H{
		"success": true,
		"message": "success",
		"data":    data,
	})
}

// Data response 200 and data
func Data(c *gin.Context, data interface{}) {
	JSON(c, gin.H{
		"success": true,
		"data":    data,
	})
}

// Created response 201 and data
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    data,
	})
}

// CreatedJSON response 201 and data
func CreatedJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

// Abort404 response 404 if not passed msg use default message
func Abort404(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": defaultMessage("Data does not exist, ensure request is correct", msg...),
	})
}

// Abort403 response 403 if not passed msg use default message
func Abort403(c *gin.Context, msg ...string) {
	c.JSON(http.StatusForbidden, gin.H{
		"message": defaultMessage("Insufficient permissions, "+
			"please make sure you have the corresponding permissions", msg...),
	})
}

// Abort500 response 500 if not passed msg use default message
func Abort500(c *gin.Context, msg ...string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": defaultMessage("Server internal error", msg...),
	})
}

// BadRequest response 400
func BadRequest(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": defaultMessage("Request parsing error, please make sure you request data is correct."+
			"upload file please using multipart header, parameters using JSON format", msg...),
		"error": err.Error(),
	})
}

// Error response 404 or 422, not passed msg use default message
func Error(c *gin.Context, err error, msg ...string) {
	logger.LogIf(err)

	// error type is database not found record
	if err == gorm.ErrRecordNotFound {
		Abort404(c)
		return
	}

	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": defaultMessage("Request process failed, please check error detail", msg...),
		"error":   err.Error(),
	})
}

// ValidationError response 422 process validation form data
func ValidationError(c *gin.Context, errors map[string][]string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": "Request validation parameters failed, please check errors.",
		"errors":  errors,
	})
}

// Unauthorized response 401
func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": defaultMessage("Request is Un Authorized please check authorization", msg...),
	})
}

// defaultMessage internal helper function
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}
	return
}
