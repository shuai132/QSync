package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf map[string]map[string]string

func GetConf() Conf {
	conf := make(Conf)
	if yamlFile, err := ioutil.ReadFile("conf.yml"); err != nil {
		log.Fatalln("配置文件读取错误", err)
	} else {
		yaml.Unmarshal(yamlFile, &conf)
	}
	return conf
}
