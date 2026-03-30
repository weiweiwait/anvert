package model

import (
	"anroid/dao"
	"time"
)

type Anthology struct {
	ID        uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID    int        `gorm:"column:user_id" json:"user_id"`
	Username  string     `gorm:"column:username" json:"username"`
	Title     string     `gorm:"column:title" json:"title"`
	Content   string     `gorm:"column:content" json:"content"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// 创建文集

func CreateAnthology(anthology *Anthology) (err error) {
	err = dao.DB.Table("anthology").Create(&anthology).Error
	return
}

// 根据用户ID查询文集（我的文集）

func GetAnthologyByUserID(userID int) ([]Anthology, error) {
	var list []Anthology
	if err := dao.DB.Debug().Table("anthology").Where("user_id=?", userID).Order("created_at DESC").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// AnthologyWithAvatar 带头像的文集结构体
type AnthologyWithAvatar struct {
	ID        uint       `gorm:"column:id" json:"id"`
	UserID    int        `gorm:"column:user_id" json:"user_id"`
	Username  string     `gorm:"column:username" json:"username"`
	Title     string     `gorm:"column:title" json:"title"`
	Content   string     `gorm:"column:content" json:"content"`
	AvatarUrl string     `gorm:"column:avatar_url" json:"avatar_url"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// 查询所有文集（诗友圈），JOIN userss表通过user_id获取头像

func GetAllAnthology() ([]AnthologyWithAvatar, error) {
	var list []AnthologyWithAvatar
	if err := dao.DB.Debug().
		Table("anthology").
		Select("anthology.*, IFNULL(userss.avatar_url, '') as avatar_url").
		Joins("LEFT JOIN userss ON anthology.user_id = userss.id").
		Order("anthology.created_at DESC").
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// 更新文集

func UpdateAnthology(id int, userID int, title string, content string) (err error) {
	now := time.Now()
	err = dao.DB.Table("anthology").Where("id=? AND user_id=?", id, userID).Updates(map[string]interface{}{
		"title":      title,
		"content":    content,
		"updated_at": now,
	}).Error
	return
}

// 删除文集（校验用户归属）

func DeleteAnthology(id int, userID int) (err error) {
	err = dao.DB.Table("anthology").Where("id=? AND user_id=?", id, userID).Delete(&Anthology{}).Error
	return
}
