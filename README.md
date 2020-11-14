<!--
 * @Author: ming
 * @LastEditors: ming
 * @Date: 2020-11-14 21:12:15
 * @LastEditTime: 2020-11-14 22:42:44
 * @FilePath: /zero-demo/README.md
 * @Description: Description
-->

# 前提

- 安装了golang
- 安装并启动docker,安装了docker-compose
- 切换到了该工程目录

# 用docker启动etcd\mysql\redis

```sh
docker-compose up -d
```

# 编译程序

## 编译api gateway

```sh
GO111MODULE=on go build -o api/_api -mod=vendor api/*.go  
```

## 编译add服务 - 写操作

```sh
GO111MODULE=on go build -o rpc/add/_add -mod=vendor rpc/add/*.go
```

## 编译check服务 - 读操作

```sh
GO111MODULE=on go build -o  rpc/check/_check -mod=vendor rpc/check/*.go
```

# 效果演示

- 创建表

```sh
docker-compose exec -T db mysql gozero < rpc/model/book.sql
```

- 启动add

```sh
rpc/add/_add -f rpc/add/etc/add.yaml > /dev/null 2>&1 &
```

- 启动check

```sh
rpc/check/_check -f rpc/check/etc/check.yaml > /dev/null 2>&1 &
```

- 启动api gateway

```sh
api/_api -f api/etc/bookstore-api.yaml > /dev/null 2>&1 &
```

- 调用服务

```sh
curl -i "http://localhost:8888/add?book=go-zero&price=10"
```

  返回如下：

  ```http
  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Thu, 03 Sep 2020 09:42:13 GMT
  Content-Length: 11
  
  {"ok":true}
  ```

```sh
curl -i "http://localhost:8888/check?book=go-zero"
```

 返回如下：

  ```http
  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Thu, 03 Sep 2020 09:47:34 GMT
  Content-Length: 25
  
  {"found":true,"price":10}
  ```

- 过载处理
 如果add服务过载了,开一台新机器,启动新的add服务注册到etcd.自动开始均衡负载.
 或者直接部署到k8s,自动伸缩.
 演示效果就修改一下 rpc/add/etc/add.yaml 端口,保存成新的配置文件,启动一个新服务.

```sh
rpc/add/_add -f rpc/add/etc/add2.yaml > /dev/null 2>&1 &
```

# 总结

该项目使用docker,快速搭建运行环境,以流畅体验go-zero.
结合科学上网,体验更佳.因为docker启动时,会自动拉取国外服务器上的镜像.

# 感谢

感谢go-zero 作者,开源了这一优秀的项目.
关于demo的实现过程,请参考官方git [链接](https://github.com/tal-tech/zero-doc/blob/main/docs/frame/bookstore.md),此处不赘述了.
