package model

import (
	"anroid/dao"
	"time"
)

type Note struct {
	ID        uint      `gorm:"id" json:"id"`
	UserID    int       `gorm:"user_id" json:"user_id"`
	UserEmail string    `gorm:"user_email" json:"user_email"`
	Title     string    `gorm:"title" json:"title"`
	Content   string    `gorm:"content" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// 创建笔记

func CreateNote(note *Note) (err error) {
	err = dao.DB.Table("note").Create(&note).Error
	return
}

// 根据用户ID查询所有笔记

func GetNotesByUserID(userID int) ([]Note, error) {
	var notes []Note
	if err := dao.DB.Debug().Table("note").Where("user_id=?", userID).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

// 根据笔记ID更新笔记

func UpdateNote(id int, userID int, title string, content string) (err error) {
	err = dao.DB.Table("note").Where("id=? AND user_id=?", id, userID).Updates(map[string]interface{}{
		"title":      title,
		"content":    content,
		"updated_at": time.Now(),
	}).Error
	return
}

// 根据笔记ID删除笔记（校验用户归属）

func DeleteNote(id int, userID int) (err error) {
	err = dao.DB.Table("note").Where("id=? AND user_id=?", id, userID).Delete(&Note{}).Error
	return
}
