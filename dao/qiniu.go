package dao

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"time"
)

var (
	AccessKey   = "dUiCN_eHnmhjFX04OUY28HE_WOrqQ_qhuX719blz"
	SecretKey   = "rSuv7jb6BcmNtdPF38vd0378NWfaJcs5u6Y9kMB8"
	Bucket      = "artverse1"
	QiNiuServer = "http://tchzkys0t.hn-bkt.clouddn.com/"
)

// 上传文件到七牛云，返回访问URL

func UploadToQiNiu(fileData []byte, fileName string) (string, error) {
	mac := qbox.NewMac(AccessKey, SecretKey)

	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	upToken := putPolicy.UploadToken(mac)

	// 自动根据 Bucket 检测所在区域
	zone, err := storage.GetZone(AccessKey, Bucket)
	if err != nil {
		fmt.Println("[七牛云] 获取区域失败:", err)
		return "", err
	}

	cfg := storage.Config{
		Zone:          zone,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 生成唯一文件名
	key := fmt.Sprintf("image/%d_%s", time.Now().UnixNano(), fileName)

	dataLen := int64(len(fileData))
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(fileData), dataLen, nil)
	if err != nil {
		fmt.Println("[七牛云] 上传失败:", err)
		return "", err
	}

	url := QiNiuServer + ret.Key
	return url, nil
}
