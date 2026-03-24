package model

import "anroid/dao"

type Signature struct {
	ID       uint   `gorm:"id"`
	Nickname string `gorm:"nickname"`
	Poem     string `gorm:"poem"`
}

// 返回所有革新前面
func GetAllSignature() ([]Signature, error) {
	var signatures []Signature
	if err := dao.DB.Debug().Table("signature").Find(&signatures).Error; err != nil {
		return nil, err
	}
	return signatures, nil
}
