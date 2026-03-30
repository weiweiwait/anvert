package server

import (
	"anroid/model"
)

// 创建笔记

func CreateNote(userID int, userEmail string, title string, content string) string {
	if userID == 0 {
		return "用户ID不能为空"
	}
	if userEmail == "" {
		return "用户邮箱不能为空"
	}
	if title == "" {
		return "笔记标题不能为空"
	}
	if content == "" {
		return "笔记内容不能为空"
	}

	note := &model.Note{
		UserID:    userID,
		UserEmail: userEmail,
		Title:     title,
		Content:   content,
	}

	err := model.CreateNote(note)
	if err != nil {
		return "创建笔记失败，请联系管理员"
	}
	return ""
}

// 查询用户所有笔记

func GetNotesByUserID(userID int) ([]model.Note, string) {
	if userID == 0 {
		return nil, "用户ID不能为空"
	}
	notes, err := model.GetNotesByUserID(userID)
	if err != nil {
		return nil, "查询笔记失败，请联系管理员"
	}
	return notes, ""
}

// 更新笔记

func UpdateNote(id int, userID int, title string, content string) string {
	if id == 0 {
		return "笔记ID不能为空"
	}
	if userID == 0 {
		return "用户ID不能为空"
	}
	if title == "" {
		return "笔记标题不能为空"
	}
	if content == "" {
		return "笔记内容不能为空"
	}

	err := model.UpdateNote(id, userID, title, content)
	if err != nil {
		return "更新笔记失败，请联系管理员"
	}
	return ""
}

// 删除笔记

func DeleteNote(id int, userID int) string {
	if id == 0 {
		return "笔记ID不能为空"
	}
	if userID == 0 {
		return "用户ID不能为空"
	}

	err := model.DeleteNote(id, userID)
	if err != nil {
		return "删除笔记失败，请联系管理员"
	}
	return ""
}
