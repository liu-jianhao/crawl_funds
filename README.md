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

![top10](https://user-images.githubusercontent.com/26483688/130358654-54778e8b-72e9-4287-97e7-d25aed81fca2.jpeg)

![held_by_funds](https://user-images.githubusercontent.com/26483688/130358662-dc6f2d4a-1e3a-4c3c-b59d-716656a1a7d1.jpeg)

