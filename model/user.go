package model

import "anroid/dao"

type User struct {
	ID       uint   `gorm:"id"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Email    string `gorm:"class"`
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
