module github.com/e421083458/gateway_demo

go 1.12

require (
	git.apache.org/thrift.git v0.13.0
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/e421083458/golang_common v1.1.1
	github.com/e421083458/gorm v1.0.1
	github.com/e421083458/grpc-proxy v0.0.0-20200322124211-7410a977f11d
	github.com/garyburd/redigo v1.6.0
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/contrib v0.0.0-20191209060500-d6e26eeaa607
	github.com/gin-gonic/gin v1.6.2
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.0
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.4
	github.com/mwitkow/grpc-proxy v0.0.0-20181017164139-0f1106ef9c76 // indirect
	github.com/pkg/errors v0.9.1
	github.com/samuel/go-zookeeper v0.0.0-20190923202752-2cc03de413da
	github.com/smartystreets/goconvey v1.6.4 // indirect
	golang.org/x/net v0.0.0-20200421231249-e086a090c8fd
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1
	google.golang.org/genproto v0.0.0-20200420144010-e5e8543f8aeb
	google.golang.org/grpc v1.30.0-dev.1
	gopkg.in/go-playground/validator.v9 v9.31.0
)

replace github.com/e421083458/gateway_demo => ./
