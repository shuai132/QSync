package qiniu

import (
	"../utils"
	"context"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"log"
)

func Upload(localFile string, bucket string, key string, accessKey string, secretKey string) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("上传成功", ret.Key, ret.Hash)
}

func UploadIfChanged(localFile string, bucket string, key string, accessKey string, secretKey string) (upload bool) {
	mac := qbox.NewMac(accessKey, secretKey)
	bucketManager := storage.NewBucketManager(mac, nil)

	fileInfo, err := bucketManager.Stat(bucket, key)
	if err != nil {
		// 没有文件等错误
		if err.Error() != "no such file or directory" {
			log.Fatalln("Stat error:", err)
		}
	} else {
		localFileHash, _ := utils.GetEtag(localFile)
		if fileInfo.Hash == localFileHash {
			log.Println("文件无变化:", localFile)
			return
		} else {
			log.Println("文件有更改:", localFile)
		}
	}

	log.Println("上传文件:", localFile)
	Upload(localFile, bucket, key, accessKey, secretKey)
	upload = true
	return
}
