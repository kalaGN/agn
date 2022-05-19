# agn
基于gin框架实现的自用api框架，学习自Gohub

    路由
    数据库
    Redis
    命令行
    代码生成（make 命令）
    验证码
    日志和错误处理
    数据库迁移
    数据填充（Faker）
    图片验证码
    分页
    授权策略
    请求验证（JSON、表单、URI Query 请求）
    分页
    限流

    缓存

    图片上传
    图片裁切


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
go run main.go make request project
```

使用redis
```
val := redis.Redis.Get("foo")
logger.Dump(val)
```
