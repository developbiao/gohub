package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/auth"
	"gohub/pkg/helpers"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// Put write data to file
func Put(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Exists check file exists
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

// FileNameWithoutExtension trim file name extension
func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func SaveUploadAvatar(c *gin.Context, file *multipart.FileHeader) (string, error) {
	var avatar string
	// Make user the directory exists, create it if it does not exist
	publicPath := "public"
	dirName := fmt.Sprintf("/upload/avatars/%s/%s/",
		app.TimenowInTimezone().Format("2006/01/02"), auth.CurrentUID(c))
	if err := os.MkdirAll(publicPath+dirName, 0755); err != nil {
		return avatar, err
	}

	// Save file
	fileName := randomNameForUploadFile(file)
	// public/upload/avatars/2023/03/21/1/xxx.png
	avatarPath := publicPath + dirName + fileName
	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return avatar, err
	}
	return avatarPath, nil
}

func randomNameForUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}
