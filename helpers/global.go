package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func CheckFileType(file *multipart.FileHeader) (string, error) {
	imgExt := []string{"png", "jpg", "jpeg", "gif", "bmp"}
	vidExt := []string{"mp4", "avi", "mkv", "mov"}

	ext := filepath.Ext(file.Filename)[1:]

	for _, val := range imgExt {
		if val == ext {
			if file.Size > 10*1024*1024 {
				return "image", fmt.Errorf("file cannot be larger than 10 mb")
			}
			return "image", nil
		}
	}

	for _, val := range vidExt {
		if val == ext {
			if file.Size > 10*1024*1024 {
				return "video", fmt.Errorf("file cannot be larger than 10 mb")
			}
			return "video", nil
		}
	}
	return "", fmt.Errorf("file type is not supported")
}

func GetUploadDir(fileName string) string {
	return "uploads/" + fileName
}

func SaveUploadedFile(c *gin.Context, uploadsFile *multipart.FileHeader) ([]byte, *os.File, error) {
	if err := c.SaveUploadedFile(uploadsFile, GetUploadDir(uploadsFile.Filename)); err != nil {
		return nil, nil, err
	}

	file, _ := os.Open(GetUploadDir(uploadsFile.Filename))

	data, err := io.ReadAll(file)
	if err != nil {
		file.Close()
		return nil, file, err
	}

	return data, file, nil
}

func ParseToJson(data any) []byte {
	json, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return json
}

func GetStage(c *gin.Context) string {
	stage, ok := c.Get("stage")
	if !ok {
		return "Development"
	}
	valid, ok := stage.(string)

	return valid
}
