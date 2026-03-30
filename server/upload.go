package server

import (
	"anroid/dao"
	"path/filepath"
	"strings"
)

// 允许的图片类型
var allowedExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
}

// 上传图片到七牛云

func UploadImage(fileData []byte, fileName string, fileSize int64) (string, string) {
	// 校验文件大小，限制5MB
	if fileSize > 5*1024*1024 {
		return "", "文件大小不能超过5MB"
	}

	// 校验文件类型
	ext := strings.ToLower(filepath.Ext(fileName))
	if !allowedExts[ext] {
		return "", "不支持的图片格式，仅支持 jpg/jpeg/png/gif/webp"
	}

	// 上传到七牛云
	url, err := dao.UploadToQiNiu(fileData, fileName)
	if err != nil {
		return "", "图片上传失败: " + err.Error()
	}

	return url, ""
}
