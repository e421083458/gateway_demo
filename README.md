# Go 微服务网关后端使用说明

大家好， 欢迎大家来到我在慕课网上的实战课程[《Go 微服务网关》](https://coding.imooc.com/class/436.html)的代码仓。这个代码仓将不仅仅包含课程的所有源代码，还将发布课程的更新相关内容，勘误信息以及计划的更多可以丰富课程的内容，如更多分享，多多练习，等等等等。

大家可以下载、运行、测试、修改。如果你发现了任何bug，或者对课程中的任何内容有意见或建议，欢迎和我联系：）

第1-8章节功能演示代码：https://github.com/e421083458/gateway_demo

项目的后端项目：https://github.com/e421083458/go_gateway

项目的前端项目：https://github.com/e421083458/go_gateway_view

项目的预览地址：http://gateway.itpp.cn:9527/

电子邮箱：e421083458@163.com

微信公众号：
![image](https://raw.githubusercontent.com/e421083458/gateway_demo/master/qrcode.jpg)

## 章节功能对应对应源码

功能点 | 地址
---|---
正在补充... | 正在补充...

## 代码帮助

### 运行后端代码

- 首先git clone 本项目

`git clone git@github.com:e421083458/gateway_demo.git`

或者 

`git clone git@github.com:e421083458/go_gateway.git`


- 确保本地环境安装了Go 1.12+版本

```
go version
go version go1.12.15 darwin/amd64
```

- 下载类库依赖

```
export GO111MODULE=on && export GOPROXY=https://goproxy.cn
cd gateway_demo
或者
cd go_gateway
go mod tidy
```

- 在相应功能文件夹下，执行 `go run main.go` 即可。

### 后端goland编辑器参考

- 用 goland 打开项目目录

- 设置 goland 支持 go mod
    - Preferences-> Go-> Go Modules（vgo）
    - 勾选 Enable Go Modules（vgo）
    - proxy 设置：https://goproxy.cn

- 在相应文件夹下的main方法中， 点击 `run go build` 即可。

### 运行前端代码

- 首先git clone 本项目

`git clone ssh://git@git.imooc.com:80/e421083458/go_gateway_view.git`

- 确保本地环境安装了nodejs

```
node -v
v11.9.0
```

- 安装node依赖包

```
cd go_gateway_view
npm install
npm install -g cnpm --registry=https://registry.npm.taobao.org
cnpm install
```

- 运行前端项目

```
npm run dev
```

### vscode编辑器设置参考

- 用 vscode 打开前端项目目录

- 安装格式化插件 ESLint、Vetur、vue-beautify

## 后端环境搭建及编辑器使用 参考文档

go环境安装介绍
http://docscn.studygolang.com/doc/install

go 基础语法学习
http://tour.studygolang.com/welcome/1

10分钟学会go mod（类库管理器）使用
https://blog.csdn.net/e421083458/article/details/89762113

goland 设置支持go mod
https://blog.csdn.net/l7l1l0l/article/details/102491573

goland 基本使用介绍
https://www.cnblogs.com/happy-king/p/9191356.html


## 前端环境搭建及编辑器使用参考文档

nodejs 安装 https://nodejs.org/zh-cn/download/

效率翻倍的 VS Code 使用指南 https://mp.weixin.qq.com/s/QpbeEgdefw2iaT8qaxkFDA