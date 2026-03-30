package server

import (
	"anroid/model"
)

// 创建文集

func CreateAnthology(userID int, username string, title string, content string) string {
	if userID == 0 {
		return "用户ID不能为空"
	}
	if username == "" {
		return "用户名不能为空"
	}
	if content == "" {
		return "文集内容不能为空"
	}

	anthology := &model.Anthology{
		UserID:   userID,
		Username: username,
		Title:    title,
		Content:  content,
	}

	err := model.CreateAnthology(anthology)
	if err != nil {
		return "创建文集失败，请联系管理员"
	}
	return ""
}

// 查询我的文集

func GetAnthologyByUserID(userID int) ([]model.Anthology, string) {
	if userID == 0 {
		return nil, "用户ID不能为空"
	}
	list, err := model.GetAnthologyByUserID(userID)
	if err != nil {
		return nil, "查询文集失败，请联系管理员"
	}
	return list, ""
}

// 查询所有文集（诗友圈）

func GetAllAnthology() ([]model.AnthologyWithAvatar, string) {
	list, err := model.GetAllAnthology()
	if err != nil {
		return nil, "查询文集失败，请联系管理员"
	}
	return list, ""
}

// 更新文集

func UpdateAnthology(id int, userID int, title string, content string) string {
	if id == 0 {
		return "文集ID不能为空"
	}
	if userID == 0 {
		return "用户ID不能为空"
	}
	if content == "" {
		return "文集内容不能为空"
	}

	err := model.UpdateAnthology(id, userID, title, content)
	if err != nil {
		return "更新文集失败，请联系管理员"
	}
	return ""
}

// 删除文集

func DeleteAnthology(id int, userID int) string {
	if id == 0 {
		return "文集ID不能为空"
	}
	if userID == 0 {
		return "用户ID不能为空"
	}

	err := model.DeleteAnthology(id, userID)
	if err != nil {
		return "删除文集失败，请联系管理员"
	}
	return ""
}
