package model

import "anroid/dao"

type Poetical struct {
	ID       uint   `gorm:"id"`
	Nickname string `gorm:"nickname"`
	Poem     string `gorm:"poem"`
}

// 返回所有诗做

func GetAllPoetical() ([]Poetical, error) {
	var poetical []Poetical
	if err := dao.DB.Debug().Table("poetical").Find(&poetical).Error; err != nil {
		return nil, err
	}
	return poetical, nil
}
