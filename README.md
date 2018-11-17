# golang 微服务学习

+ 本项目中所有可生成的部分皆已忽略，可使用 `make build` 生成

## 全部环境搭建过程

+ golang 及 常用 x包 的安装略

```bash
# grpc
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
# x/net x/text 略
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
cd $GOPATH/src/
go install google.golang.org/grpc
```

```mod
module google.golang.org/grpc

require (
	cloud.google.com/go v0.26.0 // indirect
	github.com/client9/misspell v0.3.4
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/mock v1.1.1
	github.com/golang/protobuf v1.2.0
	github.com/kisielk/gotool v1.0.0 // indirect
	golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f // indirect
	golang.org/x/sys v0.0.0-20180830151530-49385e6e1522
	golang.org/x/text v0.3.0 // indirect
	golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52
	google.golang.org/appengine v1.1.0 // indirect
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
	honnef.co/go/tools v0.0.0-20180728063816-88497007e858
)
```

```bash
# go-micro
go get -v github.com/micro/go-micro
go get -u -v github.com/micro/protobuf/{proto,protoc-gen-go}
go get -u -v github.com/micro/protoc-gen-micro
# micro工具包(可选项)
go get -v github.com/micro/micro
```

```bash
# consul 容器（go-micro 默认服务发现）
docker pull consul:1.3.0
```

## 参考

+ [主要参考博客](https://wuyin.io/)
    + [part1](https://wuyin.io/2018/05/10/microservices-part-1-introduction-and-talk-service/)
    + [part2](https://wuyin.io/2018/05/12/microservices-part-2-use-go-micro-and-dockerising/)
    + [part3](https://wuyin.io/2018/05/22/microservices-part-3-docker-compose-and-mongodb-with-orm/)
    + [part4](https://wuyin.io/2018/05/27/microservices-part-4-auth-user-by-jwt/)
    + [part5](https://wuyin.io/2018/05/28/microservices-part-5-event-brokering-with-go-micro/)
    + [part6](https://wuyin.io/2018/05/30/microservices-part-6-web-clients/)
    + [翻译者总结](https://wuyin.io/2018/06/01/microservices-summaries/)
    + [part7](https://ewanvalentine.io/microservices-in-golang-part-7/)
    + [part8](https://ewanvalentine.io/microservices-in-golang-part-8/)
    + [part9](https://ewanvalentine.io/microservices-in-golang-part-9/)
    + [part10](https://ewanvalentine.io/microservices-in-golang-part-10/)
+ [grpc 安装](https://www.jianshu.com/p/dba4c7a6d608)
+ [](https://cloud.tencent.com/developer/article/1351762)
+ [](https://wuyin.io/2018/05/02/protobuf-with-grpc-in-golang/)
