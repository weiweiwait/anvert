package server

import (
	"anroid/dao"
	"anroid/middles"
	"anroid/model"
	"strconv"
	"time"
)

func UserRegister(username string, password string, email string, code string) string {
	//连接redis
	RedisClient, err := dao.ConnectToRedis()
	// 检查Redis中键是否存在
	key := "email:" + ":" + email + ":false"
	exist, err := RedisClient.Exists(key).Result()
	if err != nil {
		return "内部错误，请联系管理员"
	}
	if exist == 0 {
		return "请先请求一封验证码邮件"
	}
	// 获取Redis中键对应的值
	result, err := RedisClient.Get(key).Result()
	if err != nil {
		return "内部错误，请联系管理员"
	}
	if result == "" {
		return "验证码失效，请重新请求"
	}
	if result == code {
		users, _ := model.FindAUserByName(username)
		if users != nil {
			return "此用户名已被注册，请更换用户名"
		}
		RedisClient.Del(key)
		// 创建新用户
		user := model.User{
			Username: username,
			Password: password,
			Email:    email,
		}

		err := model.CreateAUser(&user)
		if err != nil {
			return "内部错误，请联系管理员"
		}

		return "" // 注册成功，返回空字符串表示成功
	} else {
		return "验证码错误，请检查后再提交"
	}
}
func UserLogint(username string, password string) (string, *model.User) {
	//判断用户名是不是为空
	if username == "" {
		return "用户名不能为空", nil
	}
	// 根据用户名从数据库中获取用户信息
	user, _ := model.FindAUserByName(username)

	// 验证用户是否存在
	if user == nil {
		return "用户不存在", nil
	}
	if password != user.Password {
		return "密码不正确", nil
	}

	return "", user // 登录成功，返回空字符串和用户信息
}

// 发送验证码

func SendEmail(email string, hashAccount bool) string {
	//连接redis
	RedisClient, err := dao.ConnectToRedis()
	key := "email:" + ":" + email + ":" + strconv.FormatBool(hashAccount)
	pan, _ := RedisClient.Exists(key).Result()
	if pan == 1 {
		expire, _ := RedisClient.TTL(key).Result()
		if expire > 120*time.Second {
			return "请求频繁，请稍后再试"
		}
	}

	// 模拟查找账户
	account, _ := model.FindAUserByEmail(email)
	if hashAccount && account == nil {
		return "没有此邮件地址的账户"
	}
	if !hashAccount && account != nil {
		return "此邮箱已被其他用户注册"
	}

	// 模拟发送邮件
	result := middles.SendCode(email)
	if result == "" {
		return "邮件发送失败，请检查邮件地址是否有效"
	}

	err = RedisClient.Set(key, result, 3*time.Minute).Err()
	if err != nil {

	}

	return ""
}

// 替换用户头像：上传图片到七牛云，更新数据库头像URL

func UpdateAvatar(userID uint, fileData []byte, fileName string, fileSize int64) (string, string) {
	if userID == 0 {
		return "", "用户ID不能为空"
	}

	// 校验文件大小，限制5MB
	if fileSize > 5*1024*1024 {
		return "", "文件大小不能超过5MB"
	}

	// 上传到七牛云
	url, err := dao.UploadToQiNiu(fileData, fileName)
	if err != nil {
		return "", "头像上传失败: " + err.Error()
	}

	// 更新数据库中的头像URL
	err = model.UpdateUserAvatar(userID, url)
	if err != nil {
		return "", "更新头像失败，请联系管理员"
	}

	return url, ""
}
