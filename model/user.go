package model

import "anroid/dao"

type User struct {
	ID        uint   `gorm:"id" json:"id"`
	Username  string `gorm:"username" json:"username"`
	Password  string `gorm:"password" json:"-"`
	Email     string `gorm:"email" json:"email"`
	AvatarUrl string `gorm:"avatar_url" json:"avatar_url"`
}

// CreateATodo 创建user

func CreateAUser(user *User) (err error) {
	err = dao.DB.Table("userss").Create(&user).Error
	return
}

// 验证用户名是否重复

func FindAUserByName(username string) (user *User, err error) {
	user = new(User)
	if err = dao.DB.Debug().Table("userss").Where("username=?", username).First(user).Error; err != nil {
		return nil, err
	}
	return
}

// 验证邮箱是否被注册过

func FindAUserByEmail(email string) (user *User, err error) {
	user = new(User)
	if err = dao.DB.Debug().Table("userss").Where("email=?", email).First(user).Error; err != nil {
		return nil, err
	}
	return
}

// 更新用户头像

func UpdateUserAvatar(userID uint, avatarUrl string) (err error) {
	err = dao.DB.Table("userss").Where("id=?", userID).Update("avatar_url", avatarUrl).Error
	return
}
