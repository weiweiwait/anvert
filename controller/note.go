package controller

import (
	"anroid/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建笔记

func CreateNote(c *gin.Context) {
	var requestData struct {
		UserID    int    `form:"user_id" json:"user_id"`
		UserEmail string `form:"user_email" json:"user_email"`
		Title     string `form:"title" json:"title"`
		Content   string `form:"content" json:"content"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	result := server.CreateNote(requestData.UserID, requestData.UserEmail, requestData.Title, requestData.Content)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "创建笔记成功",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}
}

// 查询用户所有笔记

func GetNotes(c *gin.Context) {
	var requestData struct {
		UserID int `form:"user_id" json:"user_id"`
	}
	if err := c.ShouldBindQuery(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	notes, result := server.GetNotesByUserID(requestData.UserID)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": notes,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}
}

// 更新笔记

func UpdateNote(c *gin.Context) {
	var requestData struct {
		ID      int    `form:"id" json:"id"`
		UserID  int    `form:"user_id" json:"user_id"`
		Title   string `form:"title" json:"title"`
		Content string `form:"content" json:"content"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	result := server.UpdateNote(requestData.ID, requestData.UserID, requestData.Title, requestData.Content)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "更新笔记成功",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}
}

// 删除笔记

func DeleteNote(c *gin.Context) {
	var requestData struct {
		ID     int `form:"id" json:"id"`
		UserID int `form:"user_id" json:"user_id"`
	}
	// 优先从 JSON body 解析，失败则从 query 参数解析
	if err := c.ShouldBindJSON(&requestData); err != nil {
		requestData.ID = 0
		requestData.UserID = 0
	}
	if requestData.ID == 0 {
		// fallback: 从 query 参数获取
		id := c.Query("id")
		uid := c.Query("user_id")
		if id != "" {
			fmt.Sscanf(id, "%d", &requestData.ID)
		}
		if uid != "" {
			fmt.Sscanf(uid, "%d", &requestData.UserID)
		}
	}

	result := server.DeleteNote(requestData.ID, requestData.UserID)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "删除笔记成功",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}
}
