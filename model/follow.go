package model

import "anroid/dao"

type UserFollow struct {
	ID         uint   `gorm:"id" json:"id"`
	FollowerID int    `gorm:"follower_id" json:"follower_id"`
	FollowedID int    `gorm:"followed_id" json:"followed_id"`
	CreatedAt  string `gorm:"created_at" json:"created_at"`
}

// FollowerInfo 粉丝信息（用于返回给前端）
type FollowerInfo struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}

// 关注某人
func CreateFollow(followerID, followedID int) error {
	follow := UserFollow{
		FollowerID: followerID,
		FollowedID: followedID,
	}
	return dao.DB.Table("user_follow").Create(&follow).Error
}

// 取消关注
func DeleteFollow(followerID, followedID int) error {
	return dao.DB.Table("user_follow").
		Where("follower_id = ? AND followed_id = ?", followerID, followedID).
		Delete(&UserFollow{}).Error
}

// 查询是否已关注
func IsFollowing(followerID, followedID int) bool {
	var count int
	dao.DB.Table("user_follow").
		Where("follower_id = ? AND followed_id = ?", followerID, followedID).
		Count(&count)
	return count > 0
}

// 获取我的粉丝列表（谁关注了我）
func GetMyFollowers(userID int) ([]FollowerInfo, error) {
	var followers []FollowerInfo
	err := dao.DB.Table("user_follow").
		Select("userss.id as user_id, userss.username, userss.avatar_url").
		Joins("JOIN userss ON userss.id = user_follow.follower_id").
		Where("user_follow.followed_id = ?", userID).
		Order("user_follow.created_at DESC").
		Find(&followers).Error
	return followers, err
}

// 获取我关注的人列表
func GetMyFollowing(userID int) ([]FollowerInfo, error) {
	var following []FollowerInfo
	err := dao.DB.Table("user_follow").
		Select("userss.id as user_id, userss.username, userss.avatar_url").
		Joins("JOIN userss ON userss.id = user_follow.followed_id").
		Where("user_follow.follower_id = ?", userID).
		Order("user_follow.created_at DESC").
		Find(&following).Error
	return following, err
}

// 获取粉丝数量
func GetFollowerCount(userID int) int {
	var count int
	dao.DB.Table("user_follow").Where("followed_id = ?", userID).Count(&count)
	return count
}

// 获取关注数量
func GetFollowingCount(userID int) int {
	var count int
	dao.DB.Table("user_follow").Where("follower_id = ?", userID).Count(&count)
	return count
}

// 批量查询当前用户是否关注了指定用户列表
func BatchCheckFollowing(followerID int, followedIDs []int) (map[int]bool, error) {
	result := make(map[int]bool)
	if len(followedIDs) == 0 {
		return result, nil
	}
	var follows []UserFollow
	err := dao.DB.Table("user_follow").
		Where("follower_id = ? AND followed_id IN (?)", followerID, followedIDs).
		Find(&follows).Error
	if err != nil {
		return result, err
	}
	for _, f := range follows {
		result[f.FollowedID] = true
	}
	return result, nil
}
