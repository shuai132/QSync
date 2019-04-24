package qiniu

import (
	"../utils"
	"context"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"log"
)

func Upload(mac *qbox.Mac, bucket string, key string, localFile string) bool {
	log.Println("上传文件:", localFile)
	putPolicy := storage.PutPolicy {
		Scope: bucket + ":" + key,
	}
	upToken := putPolicy.UploadToken(mac)
	formUploader := storage.NewFormUploader(nil)
	ret := storage.PutRet{}

	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		log.Println("上传出错", err, localFile)
		return false
	}
	log.Println("上传成功", ret.Key, ret.Hash)
	return true
}

func UploadIfChanged(mac *qbox.Mac, bucket string, key string, localFile string) bool {
	bucketManager := storage.NewBucketManager(mac, nil)

	fileInfo, err := bucketManager.Stat(bucket, key)
	if err != nil {
		// 没有文件等错误
		if err.Error() != "no such file or directory" {
			log.Println("Stat error:", err, bucket, key, localFile)
		}
	} else {
		localFileHash, _ := utils.GetEtag(localFile)
		if fileInfo.Hash == localFileHash {
			log.Println("文件无变化:", localFile)
			return false
		} else {
			log.Println("文件有更改:", localFile)
		}
	}

	return Upload(mac, bucket, key, localFile)
}
