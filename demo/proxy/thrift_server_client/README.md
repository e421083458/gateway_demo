# 首先确保安装了thrift，安装步骤请参照
https://thrift.apache.org/docs/install/

按照相应操作系统进行安装

# 构建thrift测试server与client

- 首先编写 `thrift_gen.thrift`
- 运行IDL生成命令
`thrift --gen go thrift_gen.thrift`
- 使用生成的IDL单独构建 server 与 client 即可