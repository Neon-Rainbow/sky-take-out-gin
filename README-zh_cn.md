# 苍穹外卖 Gin框架重构版

一个使用Gin框架重构的外卖系统

技术栈:
Go + Gin + Gorm + PostgreSQL + Redis + Server-Sent Events

+ 项目框架:[Gin](https://gin-gonic.com/)
+ 数据库:[PostgreSQL](https://www.postgresql.org/)
+ 缓存:[Redis](https://redis.io/)
+ ORM工具:[Gorm](https://gorm.io/zh_CN/)
+ 项目内部通信:[Server-Sent Events](https://en.wikipedia.org/wiki/Server-sent_events)

项目实现了大多数苍穹外卖的后端接口
其中未实现的接口:
+ [ ] 微信登录(项目中使用了账号密码登录)
+ [ ] 微信支付(项目中使用了模拟支付)

该项目与苍穹外卖项目的一些接口的区别:
+ 接口使用下划线命名法,如`/api/v1/user_info`而不是`/api/v1/userInfo`
+ 接口返回的数据格式统一为:
```json
{
  "code": 0,
  "msg": "success",
  "data": {}
}
```
其中三个字段都一定会存在,但是msg和data可能为空

+ 一些参数的命名的区别,如`user_id`而不是`userId`
+ 使用了双token刷新,同时对token使用了白名单机制,将合法的token存储在了redis中,并且设置了过期时间.增加了刷新token的接口
+ 使用了Server-Sent Events实现了一个简单的消息推送功能,而不是使用WebSocket.由于项目的消息推送的时效性的要求不是很高,因此采用了更加轻量级的Server-Sent Events,同时也减少了开发成本

项目配置文件格式:
```yaml
# config.yaml

database:
  # 数据库配置,使用PostgreSQL数据库
  host: "127.0.0.1"
  port: 5432
  username: "admin"
  password: "admin"
  name: "sky_take_out"
  charset: "utf8mb4"
  parseTime: true
  loc: "UTC"


redis:
    host: "127.0.0.1"
    port: 6379
    username: "admin"
    password: "admin"
    db: 0

server:
  host: "127.0.0.1"
  port: 8080 # 服务器监听端口
  logFilePath: "./logs" #
  mode: "debug" # gin模式配置
  timeout: 3 # 服务器超时时间,单位为秒

secret:
  passwordSecret: "sky_take_out" # 密码加密秘钥
  jwtSecret: "sky_take_out" # jwt加密秘钥

log:
  level: "debug" # zap日志库等级配置
```

上传文件服务器采用了[该项目](https://github.com/Neon-Rainbow/gin-image-server),该项目运行在腾讯云服务器上.请自行修改上传文件的接口