# QSync

用于自动同步本地文件夹到七牛云空间

## Introduction
七牛云有官方工具[qshell](https://developer.qiniu.com/kodo/tools/1302/qshell) 但是不太好用

## How it works
对比本地文件和服务器文件hash，不一致或新文件则上传。

## How to use
1. 创建配置文件
```bash
cp conf.example.yml conf.yml
```
修改conf.yml中的配置

2. 编译运行
```bash
go build
./QSync
```
