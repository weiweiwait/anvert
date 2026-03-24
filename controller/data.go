package controller

import (
	"anroid/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回所有诗句

func SendPoetry(c *gin.Context) {
	allPoetry, err := model.GetAllPoems()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": allPoetry,
	})
}

// 返回所有评论

func SendComment(c *gin.Context) {
	allComment, err := model.GetAllComments()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": allComment,
	})
}

// 返回所有诗作

func SendPoetical(c *gin.Context) {
	allPoetical, err := model.GetAllPoetical()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": allPoetical,
	})
}

//返回所有个性签名

func SendSignature(c *gin.Context) {
	allSignature, err := model.GetAllSignature()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": allSignature,
	})
}
