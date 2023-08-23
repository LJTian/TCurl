# tcurl
## 介绍
### tcurl 类curl，主要增加了写入数据库的操作，还有日志延迟web散点图、折线图展示。方便直观看到访问情况、延迟情况。

## 使用
### 部署服务器
```shell
docker run -d -p 8080:8080 --name=http-server-gen  docker.io/ljtian/http-server-gen:v0.2 
```

### 部署数据库
```shell
docker run -d -p 3306:3306 --name mysql-container -e MYSQL_ROOT_PASSWORD=123456 mysql
```
### 部署客户端

```shell
./tcurl -U=http://127.0.0.1:8080/ping -S=true -T=10000 -t=1 -D="root:123456@tcp(10.12.17.231:3306)/logs?charset=utf8mb4&parseTime=True&loc=Local" -I=1 -c=5
```

### web 展示数据
```shell
./tcurl show -D="root:123456@tcp(127.0.0.1:3306)/logs?charset=utf8mb4&parseTime=True&loc=Local" -C=clientName_jK0zg
```

## 授权
Apache License Version 2.0

