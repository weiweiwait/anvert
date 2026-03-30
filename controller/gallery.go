package controller

import (
	"anroid/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// 创建画作（支持图片上传）

func CreateGallery(c *gin.Context) {
	userIDStr := c.PostForm("user_id")
	username := c.PostForm("username")
	title := c.PostForm("title")
	creator := c.PostForm("creator")
	year := c.PostForm("year")
	material := c.PostForm("material")
	size := c.PostForm("size")
	description := c.PostForm("description")

	var userID int
	if _, err := fmt.Sscanf(userIDStr, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": "用户ID格式错误",
		})
		return
	}

	// 获取上传的图片文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": "请选择要上传的画作图片",
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

	// 先上传图片到七牛云
	imageURL, uploadResult := server.UploadImage(fileData, header.Filename, header.Size)
	if uploadResult != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": uploadResult,
		})
		return
	}

	// 创建画作记录
	result := server.CreateGallery(userID, username, title, imageURL, creator, year, material, size, description)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "创建画作成功",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}
}

// 获取我的画廊

func GetMyGallery(c *gin.Context) {
	var requestData struct {
		UserID int `form:"user_id" json:"user_id"`
	}
	if err := c.ShouldBindQuery(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	list, result := server.GetGalleryByUserID(requestData.UserID)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": list,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}
}

// 获取画作详情

func GetGalleryDetail(c *gin.Context) {
	idStr := c.Query("id")
	var id int
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": "画作ID格式错误",
		})
		return
	}

	gallery, result := server.GetGalleryByID(id)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": gallery,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}
}

// 更新画作（支持重新上传图片）

func UpdateGallery(c *gin.Context) {
	userIDStr := c.PostForm("user_id")
	idStr := c.PostForm("id")
	title := c.PostForm("title")
	creator := c.PostForm("creator")
	year := c.PostForm("year")
	material := c.PostForm("material")
	size := c.PostForm("size")
	description := c.PostForm("description")

	var userID, id int
	fmt.Sscanf(userIDStr, "%d", &userID)
	fmt.Sscanf(idStr, "%d", &id)

	// 如果有新图片上传，则替换
	imageURL := ""
	file, header, err := c.Request.FormFile("file")
	if err == nil {
		defer file.Close()
		fileData, readErr := io.ReadAll(file)
		if readErr == nil {
			url, uploadResult := server.UploadImage(fileData, header.Filename, header.Size)
			if uploadResult != "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"Code":    http.StatusBadRequest,
					"Message": uploadResult,
				})
				return
			}
			imageURL = url
		}
	}

	result := server.UpdateGallery(id, userID, title, imageURL, creator, year, material, size, description)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "更新画作成功",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}
}

// 删除画作

func DeleteGallery(c *gin.Context) {
	var requestData struct {
		ID     int `form:"id" json:"id"`
		UserID int `form:"user_id" json:"user_id"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		requestData.ID = 0
		requestData.UserID = 0
	}
	if requestData.ID == 0 {
		id := c.Query("id")
		uid := c.Query("user_id")
		if id != "" {
			fmt.Sscanf(id, "%d", &requestData.ID)
		}
		if uid != "" {
			fmt.Sscanf(uid, "%d", &requestData.UserID)
		}
	}

	result := server.DeleteGallery(requestData.ID, requestData.UserID)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "删除画作成功",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}
}
