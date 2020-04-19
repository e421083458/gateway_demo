# 安装zookeeper

- 参考官方文档安装
http://zookeeper.apache.org/doc/r3.6.0/zookeeperStarted.html
- 解压缩
- 编辑 conf/zoo.cfg
```
tickTime=2000
dataDir=/var/lib/zookeeper
clientPort=2181
```
- 运行 `bin/zkServer.sh start`

