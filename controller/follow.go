package controller

import (
	"anroid/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 关注用户
func FollowUser(c *gin.Context) {
	var requestData struct {
		FollowerID int `form:"follower_id" json:"follower_id"`
		FollowedID int `form:"followed_id" json:"followed_id"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	result := server.FollowUser(requestData.FollowerID, requestData.FollowedID)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "关注成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": result,
		})
	}
}

// 取消关注
func UnfollowUser(c *gin.Context) {
	var requestData struct {
		FollowerID int `form:"follower_id" json:"follower_id"`
		FollowedID int `form:"followed_id" json:"followed_id"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	result := server.UnfollowUser(requestData.FollowerID, requestData.FollowedID)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "已取消关注",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": result,
		})
	}
}

// 获取我的粉丝列表
func GetMyFollowers(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}
	followers, err := server.GetMyFollowers(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	followerCount, followingCount := server.GetFollowCounts(userID)
	c.JSON(http.StatusOK, gin.H{
		"Code":            http.StatusOK,
		"data":            followers,
		"follower_count":  followerCount,
		"following_count": followingCount,
	})
}

// 获取我关注的人列表
func GetMyFollowing(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}
	following, err := server.GetMyFollowing(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	followerCount, followingCount := server.GetFollowCounts(userID)
	c.JSON(http.StatusOK, gin.H{
		"Code":            http.StatusOK,
		"data":            following,
		"follower_count":  followerCount,
		"following_count": followingCount,
	})
}

// 检查是否已关注（单个）
func CheckFollow(c *gin.Context) {
	followerIDStr := c.Query("follower_id")
	followedIDStr := c.Query("followed_id")
	followerID, err1 := strconv.Atoi(followerIDStr)
	followedID, err2 := strconv.Atoi(followedIDStr)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的参数"})
		return
	}
	result, err := server.BatchCheckFollowing(followerID, []int{followedID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Code":         http.StatusOK,
		"is_following": result[followedID],
	})
}

// 批量检查关注状态（用于诗友圈列表）
func BatchCheckFollow(c *gin.Context) {
	var requestData struct {
		FollowerID  int   `json:"follower_id"`
		FollowedIDs []int `json:"followed_ids"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	result, err := server.BatchCheckFollowing(requestData.FollowerID, requestData.FollowedIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Code": http.StatusOK,
		"data": result,
	})
}
