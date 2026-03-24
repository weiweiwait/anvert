package model

import "anroid/dao"

type Comment struct {
	ID       uint   `gorm:"id"`
	Nickname string `gorm:"nickname"`
	Poem     string `gorm:"poem"`
}

// 返回所有评论

func GetAllComments() ([]Comment, error) {
	var comments []Comment
	if err := dao.DB.Debug().Table("comment").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// CreateATodo 创建comment

func CreatePoem(comment *Comment) (err error) {
	err = dao.DB.Table("comment").Create(&comment).Error
	return
}
