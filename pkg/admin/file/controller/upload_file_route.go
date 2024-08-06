package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// FileRoutes 文件路由
// @Summary 文件路由
// 该函数用于处理文件上传
func FileRoutes(routes *gin.RouterGroup) {
	// 上传图片的接口
	routes.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("image")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		tempFilePath := "./" + header.Filename
		out, err := os.Create(tempFilePath)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("create temp file err: %s", err.Error()))
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("write temp file err: %s", err.Error()))
			return
		}

		resp, err := uploadFileToServer(tempFilePath, "http://124.223.10.155:8080/image-server/upload")
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("upload image err: %s", err.Error()))
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			c.String(http.StatusInternalServerError, fmt.Sprintf("upload image failed, status code: %d", resp.StatusCode))
			return
		}

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("decode response err: %s", err.Error()))
			return
		}

		// 修改 filepath 字段
		if filepath, ok := result["filepath"].(string); ok {
			result["filepath"] = "http://124.223.10.155:8080/image-server/image/" + filepath[len("uploads/"):]
		}

		c.JSON(http.StatusOK, result)

		os.Remove(tempFilePath)
	})
}

// UploadFile 上传文件
// @Summary 上传文件
// @Param file formData file true "文件"
// @Router /api/v1/common/upload [post]
func uploadFileToServer(filePath string, targetURL string) (*http.Response, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	writer.Close()

	request, err := http.NewRequest("POST", targetURL, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	return client.Do(request)
}
