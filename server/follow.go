package server

import "anroid/model"

// 关注某人
func FollowUser(followerID, followedID int) string {
	if followerID == followedID {
		return "不能关注自己"
	}
	if model.IsFollowing(followerID, followedID) {
		return "已经关注过了"
	}
	err := model.CreateFollow(followerID, followedID)
	if err != nil {
		return "关注失败: " + err.Error()
	}
	return ""
}

// 取消关注
func UnfollowUser(followerID, followedID int) string {
	if !model.IsFollowing(followerID, followedID) {
		return "尚未关注该用户"
	}
	err := model.DeleteFollow(followerID, followedID)
	if err != nil {
		return "取消关注失败: " + err.Error()
	}
	return ""
}

// 获取我的粉丝列表
func GetMyFollowers(userID int) ([]model.FollowerInfo, error) {
	return model.GetMyFollowers(userID)
}

// 获取我关注的人
func GetMyFollowing(userID int) ([]model.FollowerInfo, error) {
	return model.GetMyFollowing(userID)
}

// 获取关注/粉丝数量
func GetFollowCounts(userID int) (int, int) {
	followerCount := model.GetFollowerCount(userID)
	followingCount := model.GetFollowingCount(userID)
	return followerCount, followingCount
}

// 批量查询关注状态
func BatchCheckFollowing(followerID int, followedIDs []int) (map[int]bool, error) {
	return model.BatchCheckFollowing(followerID, followedIDs)
}
