# Cinema - 电影院订票管理

# 概述

使用 `go-micro` 开发的微服务电影院订票管理案例

目录结构：
- **movie** - 电影管理服务
- **proto** - `grpc` 协议文件目录
- **projection** - 电影票管理服务
- **user** - 用户管理服务

# 特性
- **微服务框架**： [go-micro](https://github.com/micro/go-micro) 微服务的基础框架
- **服务注册/发现**：[consul](https://github.com/hashicorp/consul) 分布式、高可用的服务注册/发现
- **消息中间件**：[kafka](https://github.com/apache/kafka) 发布/订阅、异步通知
- **api网关**：[gin](https://github.com/gin-gonic/gin) Api统一管理、鉴权
- **JWT鉴权**：[jwt-go](https://github.com/dgrijalva/jwt-go) - JSON Web Token
- **api文档**：`swagger` api 文档
- **rpc协议**：[grpc-go](https://github.com/grpc/grpc-go) 服务间的通信协议
- **错误报告**：[sentry](https://github.com/getsentry/raven-go) 警报系统
- **缓存**：`redis` 缓存客户端使用[go-redis](https://github.com/go-redis/redis)
- **DB**：`Mongodb` 客户端使用[mongo-go-driver](https://github.com/mongodb/mongo-go-driver)
- **部署**：[docker-compose](https://github.com/docker/compose) 部署

# 依赖
- go1.11+
- docker-compose
# 功能完善中...


