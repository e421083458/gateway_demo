## 192.168.3.4是你当前机器的局域网IP


安装地址：
https://github.com/mlabouardy/hystrix-dashboard-docker

需要提前安装一下docker环境，然后执行下载操作：

$git clone git@github.com:mlabouardy/hystrix-dashboard-docker.git


切换到下载目录：
$cd hystrix-dashboard-docker

执行服务启动
$docker run -d -p 8080:9002 --name hystrix-dashboard mlabouardy/hystrix-dashboard:latest

此时就可以打本地的hystrix dashboard地址了：

http://192.168.3.4:8080
=====================================


在以下地址下面运行测试用例，你会启动一个Stream server

https://github.com/e421083458/gateway_demo/blob/master/demo/proxy/circuit_breaker/main_test.go

Stream server 地址： 

http://192.168.3.4:8074/

=====================================

地址一正常情况下会有一个输入框，你把地址二贴上去就可以查到效果了。
