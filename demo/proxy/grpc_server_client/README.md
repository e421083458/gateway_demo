# 关于grpc方面功能演示为了简便期间，统一使用非加密方式

# 首先确保安装了grpc 安装，安装步骤请参照
https://github.com/grpc/grpc-go

## 安装步骤
### go mod安装方式
- 开启 go mod `export GO111MODULE=on`
- 开启代理 go mod `export GOPROXY=https://goproxy.io`
- 安装grpc `go get -u google.golang.org/grpc`
- 安装proto `go get -u github.com/golang/protobuf/proto`
- 安装protoc-gen-go `go get -u github.com/golang/protobuf/protoc-gen-go`


### 遇到命令不存在 command not found: protoc

```
缺少protobuf预先安装导致的问题。 protoc-gen-go相当于只是protobuf的一个语言支持。

安装protobuf参照资料
mac 下   brew install protobuf
windows 下 https://blog.csdn.net/qq_41185868/article/details/82882206
linux 下 https://blog.csdn.net/sszzyzzy/article/details/89946075
```

### 遇到 I/O Timeout Errors

```
$ go get -u google.golang.org/grpc
package google.golang.org/grpc: unrecognized import path "google.golang.org/grpc" (https fetch: Get https://google.golang.org/grpc?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
```

- 要么翻墙解决问题
- 要么设置 `export GOPROXY=https://goproxy.io`

### 遇到 permission denied

```
go get -u google.golang.org/grpc
go get github.com/golang/protobuf/protoc-gen-go: open /usr/local/go/bin/protoc-gen-go: permission denied
```

- 要么给目录增加读写权限 `chmod -R 777 /usr/local/go/`
- 要么设置重新设置一下GOROOT地址指向非系统目录 
比如 mac 下面 
vim ~/.bashrc
```
#export GOROOT=/usr/local/go
#export GOBIN=/usr/local/go/bin
#export LGOBIN=/usr/local/go/bin
export GOROOT=/Users/niuyufu/go_1.12
export GOBIN=/Users/niuyufu/go_1.12/bin
export LGOBIN=/Users/niuyufu/go_1.12/bin
 ```

# 构建grpc测试server与client

- 首先编写 `echo.proto`
- 运行IDL生成命令，（如：遇到命令不存在 command not found: protoc 请参照上文帮助）

`protoc -I . --go_out=plugins=grpc:proto ./echo.proto`
- 使用生成的IDL单独构建 server 与 client 即可


# 构建grpc-gateway 测试服务端让服务器支持http

## 安装参考
`https://github.com/grpc-ecosystem/grpc-gateway`
- 开启 go mod `export GO111MODULE=on`
- 开启代理 go mod `export GOPROXY=https://goproxy.io`
- 执行安装命令

```
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go install  github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go install github.com/golang/protobuf/protoc-gen-go
```

## 构建grpc-gateway 测试服务端

- 编写 `echo-gateway.proto`
- 运行IDL生成命令
```
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:proto echo-gateway.proto
```
- 删除 proto/echo.pb.go 防止结构体冲突
`rm proto/echo.pb.go`
- 运行gateway生成命令
```
protoc -I/usr/local/include -I. -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:proto echo-gateway.proto
```
- 使用生成的IDL单独构建 server
- 使用浏览器测试 server
```
curl 'http://127.0.0.1:8081/v1/example/echo' -d '{"message":"11222222"}'
```
