package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf map[string]map[string]string

func GetConf(filePath string) Conf {
	conf := make(Conf)
	if yamlFile, err := ioutil.ReadFile(filePath); err != nil {
		log.Fatalln("配置文件读取错误", err)
	} else {
		_ = yaml.Unmarshal(yamlFile, &conf)
	}
	return conf
}
