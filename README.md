<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [功能对应对应源码](#%E5%8A%9F%E8%83%BD%E5%AF%B9%E5%BA%94%E5%AF%B9%E5%BA%94%E6%BA%90%E7%A0%81)
- [Go 微服务网关代码使用说明](#go-%E5%BE%AE%E6%9C%8D%E5%8A%A1%E7%BD%91%E5%85%B3%E4%BB%A3%E7%A0%81%E4%BD%BF%E7%94%A8%E8%AF%B4%E6%98%8E)
  - [代码帮助](#%E4%BB%A3%E7%A0%81%E5%B8%AE%E5%8A%A9)
    - [运行后端代码](#%E8%BF%90%E8%A1%8C%E5%90%8E%E7%AB%AF%E4%BB%A3%E7%A0%81)
    - [后端goland编辑器参考](#%E5%90%8E%E7%AB%AFgoland%E7%BC%96%E8%BE%91%E5%99%A8%E5%8F%82%E8%80%83)
    - [运行前端代码](#%E8%BF%90%E8%A1%8C%E5%89%8D%E7%AB%AF%E4%BB%A3%E7%A0%81)
    - [vscode编辑器设置参考](#vscode%E7%BC%96%E8%BE%91%E5%99%A8%E8%AE%BE%E7%BD%AE%E5%8F%82%E8%80%83)
  - [后端环境搭建及编辑器使用 参考文档](#%E5%90%8E%E7%AB%AF%E7%8E%AF%E5%A2%83%E6%90%AD%E5%BB%BA%E5%8F%8A%E7%BC%96%E8%BE%91%E5%99%A8%E4%BD%BF%E7%94%A8-%E5%8F%82%E8%80%83%E6%96%87%E6%A1%A3)
  - [前端环境搭建及编辑器使用参考文档](#%E5%89%8D%E7%AB%AF%E7%8E%AF%E5%A2%83%E6%90%AD%E5%BB%BA%E5%8F%8A%E7%BC%96%E8%BE%91%E5%99%A8%E4%BD%BF%E7%94%A8%E5%8F%82%E8%80%83%E6%96%87%E6%A1%A3)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# 功能对应对应源码

功能点| 源码地址
---|---
熔断器| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/circuit_breaker)
单机流量统计| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/flow_count)
浏览器正常代理| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/forward_proxy)
grpc反向代理| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/grpc_reverse_proxy)
grpc反向代理整合中间件| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/grpc_reverse_proxy_advance)
grpc反向代理整合负载均衡器| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/grpc_reverse_proxy_lb)
grpc测试服务器、客户端、grpc-gateway| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/grpc_server_client)
负载均衡器| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/load_balance)
负载均衡主动探测| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/load_balance_client_discovery)
负载均衡服务发现| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/load_balance_server_discovery)
中间件实现| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/middleware)
观察者模式| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/observer)
限流器| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/rate_limiter)
测试下游服务器| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/real_server)
测试下游服务器+服务注册| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/real_server_register)
分布式流量统计| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/redis_flow_count)
http反向代理实现| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/reverse_proxy)
http反向代理简单版| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/reverse_proxy_simple)
http2反向代理| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/reverse_proxy_http2)
https反向代理| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/reverse_proxy_https)
http反向代理基本功能| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/reverse_proxy_base)
http反向代理权限校验| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/security_check)
tcp代理服务器实现| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/tcp_proxy)
thrift服务器与客户端| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/thrift_server_client)
websocket代理服务器| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/websocket)
websocket代理服务器| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/websocket)
zookeeper基本使用| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/proxy/zookeeper)
==基础==| ===基础===
函数是一等公民| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/base/functional)
http客户端| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/base/http_client)
http服务端| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/base/http_server)
tcp客户端| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/base/tcp_client)
tcp代理| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/base/tcp_proxy)
tcp服务器| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/base/tcp_server)
udp客户端| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/base/udp_client)
udp服务端| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/base/udp_server)
自定义协议获取完整报文| [源代码](https://github.com/e421083458/gateway_demo/tree/master/demo/base/unpack)

其他正在补充...| [源代码](其他正在补充...


# Go 微服务网关代码使用说明

这是慕课网上的实战课程[《Go 微服务网关》](https://coding.imooc.com/class/436.html)的代码仓。这个代码仓将不仅仅包含课程的所有源代码，还将发布课程的更新相关内容，勘误信息以及计划的更多可以丰富课程的内容，如更多分享，多多练习，等等等等。

大家可以下载、运行、测试、修改。如果你发现了任何bug，或者对课程中的任何内容有意见或建议，欢迎和我联系：）

第1-8章节功能演示代码：https://github.com/e421083458/gateway_demo

完整后端项目：https://github.com/e421083458/go_gateway

完整前端项目：https://github.com/e421083458/go_gateway_view

项目的预览地址：http://gateway.itpp.cn:9527/

电子邮箱：e421083458@163.com

微信公众号：

![image](http://chuantu.xyz/t6/731/1587960911x3030586988.jpg)

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