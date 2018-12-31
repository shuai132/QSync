package main

import (
	"./qiniu"
	"./utils"
	"github.com/qiniu/api.v7/auth/qbox"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	config := GetConf(utils.GetFullPath("./conf.yml"))

	accessKey := config["qiniu"]["ACCESS_KEY"]
	secretKey := config["qiniu"]["SECRET_KEY"]
	bucketName := config["qiniu"]["BUCKET_NAME"]
	mac := qbox.NewMac(accessKey, secretKey)

	resDir := config["path"]["RES_DIR"]
	regexFileName, _ := regexp.Compile(config["path"]["PATTERN_FILENAME"])
	regexFilePath, _ := regexp.Compile(config["path"]["PATTERN_FILEPATH"])
	regexIgnoreFile, _ := regexp.Compile(config["path"]["PATTERN_IGNOREFILE"])

	log.Println("本地资源目录:", resDir)

	ch := make(chan bool)
	goroutineCount := 0
	filepath.Walk(resDir, func(fullPath string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		fileName := filepath.Base(fullPath)
		relativePath := fullPath[len(resDir)+1:]

		if !regexFileName.MatchString(fileName) {
			return nil
		}
		if !regexFilePath.MatchString(relativePath) {
			return nil
		}
		if regexIgnoreFile.MatchString(fileName) {
			return nil
		}

		go func() {
			ch <- qiniu.UploadIfChanged(mac, bucketName, relativePath, fullPath)
		}()

		goroutineCount++
		return nil
	})

	uploadCount := 0
	for ; goroutineCount > 0; goroutineCount-- {
		if <-ch {
			uploadCount++
		}
	}
	log.Println("同步完成: 上传文件数量:", uploadCount)
}
