package controller

import (
	"anroid/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建文集

func CreateAnthology(c *gin.Context) {
	var requestData struct {
		UserID   int    `form:"user_id" json:"user_id"`
		Username string `form:"username" json:"username"`
		Title    string `form:"title" json:"title"`
		Content  string `form:"content" json:"content"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	result := server.CreateAnthology(requestData.UserID, requestData.Username, requestData.Title, requestData.Content)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "创建文集成功",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}
}

// 获取我的文集

func GetMyAnthology(c *gin.Context) {
	var requestData struct {
		UserID int `form:"user_id" json:"user_id"`
	}
	if err := c.ShouldBindQuery(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	list, result := server.GetAnthologyByUserID(requestData.UserID)
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

// 获取所有文集（诗友圈）

func GetAllAnthology(c *gin.Context) {
	list, result := server.GetAllAnthology()
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

// 更新文集

func UpdateAnthology(c *gin.Context) {
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

	result := server.UpdateAnthology(requestData.ID, requestData.UserID, requestData.Title, requestData.Content)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "更新文集成功",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}
}

// 删除文集

func DeleteAnthology(c *gin.Context) {
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

	result := server.DeleteAnthology(requestData.ID, requestData.UserID)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "删除文集成功",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}
}
