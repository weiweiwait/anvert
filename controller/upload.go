package controller

import (
	"anroid/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// 上传图片

func UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": "请选择要上传的图片",
		})
		return
	}
	defer file.Close()

	// 读取文件内容
	fileData, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Code":    http.StatusInternalServerError,
			"Message": "读取文件失败",
		})
		return
	}

	url, result := server.UploadImage(fileData, header.Filename, header.Size)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "上传成功",
			"url":     url,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": result,
		})
	}
}

// 替换用户头像

func UpdateAvatar(c *gin.Context) {
	// 获取用户ID
	userIDStr := c.PostForm("user_id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": "用户ID不能为空",
		})
		return
	}

	var userID uint
	if _, err := fmt.Sscanf(userIDStr, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": "用户ID格式错误",
		})
		return
	}

	// 获取上传的图片
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": "请选择要上传的头像图片",
		})
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Code":    http.StatusInternalServerError,
			"Message": "读取文件失败",
		})
		return
	}

	url, result := server.UpdateAvatar(userID, fileData, header.Filename, header.Size)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":       http.StatusOK,
			"Message":    "头像更新成功",
			"avatar_url": url,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": result,
		})
	}
}
