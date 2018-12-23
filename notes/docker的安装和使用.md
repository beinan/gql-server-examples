# Docker的安装和使用

## 安装Docker

## 启动一个Docker容器

## 涉及到的一些Docker命令

### 列出docker容器
`docker ps`

### 登录到docker容器中
`docker exec -it {容器名字} /bin/bash`
例如:
`docker exec -it gqlserverexamples_gateway_1 /bin/bash`

## docker-compose的使用

### 启动
`docker-compose up`

### 关闭
`docker-compose down`

### 水平扩展
`docker-compose up --scale user-service=4`

### 查看日志
`docker-compose logs -f`


