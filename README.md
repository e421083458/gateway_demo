# Go 微服务网关配套后端代码

前端项目：https://github.com/e421083458/gateway_demo_view

## 代码帮助

### 项目目录

```
├── README.md
├── demo            示例代码
├── go.mod          go mod文件
├── project         项目代码
└── proxy           公共类库
```

### 使用goland 运行后端代码

- 首先git clone 本项目

`git clone git@github.com:e421083458/gateway_demo.git`

- 确保本地环境安装了Go 1.12+版本

```
go version
go version go1.12.15 darwin/amd64
```

- 下载类库依赖

```
export GO111MODULE=on && export GOPROXY=https://goproxy.cn
cd gateway_demo
go mod tidy
```

- 用 goland 打开项目目录

- 设置 goland 支持 go mod
    - Preferences-> Go-> Go Modules（vgo）
    - 勾选 Enable Go Modules（vgo）
    - proxy 设置：https://goproxy.cn

- 在相应文件夹下的main方法中， 点击 `run go build` 即可。

## 环境搭建及编辑器使用 参考文档

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
