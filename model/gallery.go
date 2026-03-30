package model

import (
	"anroid/dao"
	"time"
)

type Gallery struct {
	ID          uint       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID      int        `gorm:"column:user_id" json:"user_id"`
	Username    string     `gorm:"column:username" json:"username"`
	Title       string     `gorm:"column:title" json:"title"`
	ImageURL    string     `gorm:"column:image_url" json:"image_url"`
	Creator     string     `gorm:"column:creator" json:"creator"`
	Year        string     `gorm:"column:year" json:"year"`
	Material    string     `gorm:"column:material" json:"material"`
	Size        string     `gorm:"column:size" json:"size"`
	Description string     `gorm:"column:description" json:"description"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// 创建画作

func CreateGallery(gallery *Gallery) error {
	return dao.DB.Table("gallery").Create(gallery).Error
}

// 根据用户ID查询画作（我的画廊）

func GetGalleryByUserID(userID int) ([]Gallery, error) {
	var list []Gallery
	if err := dao.DB.Debug().Table("gallery").Where("user_id=?", userID).Order("created_at DESC").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// 根据ID查询单个画作详情

func GetGalleryByID(id int) (*Gallery, error) {
	var gallery Gallery
	if err := dao.DB.Debug().Table("gallery").Where("id=?", id).First(&gallery).Error; err != nil {
		return nil, err
	}
	return &gallery, nil
}

// 更新画作

func UpdateGallery(id int, userID int, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()
	return dao.DB.Table("gallery").Where("id=? AND user_id=?", id, userID).Updates(updates).Error
}

// 删除画作（校验用户归属）

func DeleteGallery(id int, userID int) error {
	return dao.DB.Table("gallery").Where("id=? AND user_id=?", id, userID).Delete(&Gallery{}).Error
}
