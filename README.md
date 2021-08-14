# golang 实现基金爬取

## 运行
该程序依赖MySQL，所以运行前需要安装MySQL

将`dal`目录下的连接语句的用户名和密码替换一下：
```go
// {user}:{password}@tcp({ip:port})/{database_name}?charset=utf8mb4&parseTime=True&loc=Local
dsn := "root:root_password@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
```

运行：
```shell
go run .
```
