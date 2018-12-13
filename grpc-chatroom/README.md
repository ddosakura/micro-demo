# 基于 gRPC 和 jwt 的简易聊天室

## 依赖

### 后端

+ grpc
+ github.com/dgrijalva/jwt-go
+ github.com/satori/go.uuid

### web 前端

+ [google-protobuf](https://www.npmjs.com/package/google-protobuf)
+ [grpc-web](https://www.npmjs.com/package/grpc-web)

#### 与 react 集成（我的方案，需要 eject 之后）

加 `.eslintignore` 忽略掉所有 proto 产物

#### 与 react 集成（网上的方案，需要 eject 之后）

+ [方案来源](http://www.cnblogs.com/xiaoxiaopao/p/9955731.html)

##### eslintConfig 配置（全局变量）

照理应该是这样的：

```json
{
  ...,
  "eslintConfig": {
    "extends": "react-app",
    "globals": {
      "COMPILED": true,
      "proto": true
    }
  },
  ...
}
```

但这样不知道为啥也可以，

```json
{
  ...,
  "eslintConfig": {
    "extends": "react-app",
    "globals": {
      "COMPILED": false,
      "proto": false
    }
  },
  ...
}
```

##### CommonJS 的问题

+ 改变 package.json 中 babel plugin 配置，把 transform-runtime 配置去掉，使其兼容 common.js 的 module.export 语法
