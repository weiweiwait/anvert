package server

import (
	"anroid/model"
)

// 创建画作

func CreateGallery(userID int, username string, title string, imageURL string, creator string, year string, material string, size string, description string) string {
	if userID == 0 {
		return "用户ID不能为空"
	}
	if title == "" {
		return "画作名称不能为空"
	}
	if imageURL == "" {
		return "请上传画作图片"
	}

	gallery := &model.Gallery{
		UserID:      userID,
		Username:    username,
		Title:       title,
		ImageURL:    imageURL,
		Creator:     creator,
		Year:        year,
		Material:    material,
		Size:        size,
		Description: description,
	}

	err := model.CreateGallery(gallery)
	if err != nil {
		return "创建画作失败，请联系管理员"
	}
	return ""
}

// 查询我的画廊

func GetGalleryByUserID(userID int) ([]model.Gallery, string) {
	if userID == 0 {
		return nil, "用户ID不能为空"
	}
	list, err := model.GetGalleryByUserID(userID)
	if err != nil {
		return nil, "查询画廊失败，请联系管理员"
	}
	return list, ""
}

// 查询画作详情

func GetGalleryByID(id int) (*model.Gallery, string) {
	if id == 0 {
		return nil, "画作ID不能为空"
	}
	gallery, err := model.GetGalleryByID(id)
	if err != nil {
		return nil, "查询画作详情失败"
	}
	return gallery, ""
}

// 更新画作

func UpdateGallery(id int, userID int, title string, imageURL string, creator string, year string, material string, size string, description string) string {
	if id == 0 {
		return "画作ID不能为空"
	}
	if userID == 0 {
		return "用户ID不能为空"
	}

	updates := map[string]interface{}{}
	if title != "" {
		updates["title"] = title
	}
	if imageURL != "" {
		updates["image_url"] = imageURL
	}
	// 这些字段允许置空，所以始终更新
	updates["creator"] = creator
	updates["year"] = year
	updates["material"] = material
	updates["size"] = size
	updates["description"] = description

	err := model.UpdateGallery(id, userID, updates)
	if err != nil {
		return "更新画作失败，请联系管理员"
	}
	return ""
}

// 删除画作

func DeleteGallery(id int, userID int) string {
	if id == 0 {
		return "画作ID不能为空"
	}
	if userID == 0 {
		return "用户ID不能为空"
	}

	err := model.DeleteGallery(id, userID)
	if err != nil {
		return "删除画作失败，请联系管理员"
	}
	return ""
}
