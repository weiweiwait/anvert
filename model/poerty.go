package model

import "anroid/dao"

type Poem struct {
	ID       uint   `gorm:"id"`
	Nickname string `gorm:"nickname"`
	Poem     string `gorm:"poem"`
}

// 返回所有诗句

func GetAllPoems() ([]Poem, error) {
	var poems []Poem
	if err := dao.DB.Debug().Table("poetry").Find(&poems).Error; err != nil {
		return nil, err
	}
	return poems, nil
}
