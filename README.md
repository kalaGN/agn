# agn
基于gin框架实现的api框架，学习自Gohub，不包含具体业务，可用于接口快速开发。

    路由
    数据库
    Redis
    缓存
    命令行
    代码生成（make 命令）
    验证码
    日志和错误处理
    数据库迁移
    图片验证码
    请求验证（JSON、表单、URI Query 请求）
    限流
    数据填充（Faker）
    分页


    图片上传
    图片裁切
    授权策略


## 使用流程

下载
```
git clone https://github.com/kalaGN/agn.git
```
修改配置文件
```
cp .env.example .env && vim .env
```
生成控制器
```
go run main.go make apicontroller v1/User
```
生成模型
```
go run main.go make model User
```

生成请求校验文件

```
go run main.go make request User
```

使用redis
```
val := redis.Redis.Get("foo")
logger.Dump(val)
```
迁移命令-新增
```
go run main.go make migration add_users_table
```
迁移命令-建表

```
go run main.go migrate up
```

缓存-清除 REDIS_CACHE_DB 对应redis db
```
go run main.go cache clear
```


缓存-清除单个key 比如 agn:cache:key1
```
go run main.go cache forget --key=key1
```

