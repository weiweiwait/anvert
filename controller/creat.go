package controller

import (
	"anroid/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//添加评论

func CreatComments(c *gin.Context) {
	var requestData struct {
		Nickname string `form:"nickname" json:"nickname"`
		Poem     string `form:"poem" json:"poem"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	comment := &model.Comment{
		Nickname: requestData.Nickname,
		Poem:     requestData.Poem,
	}

	err := model.CreatePoem(comment)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": "内部错误",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "评论成功",
		})
	}
}
