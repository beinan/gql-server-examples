# 安装Golang 

## 下载安装
* 官网：
https://golang.org/dl/

# 编写Go代码

## GOPATH

把下面的配置加入~/.zshrc
```zsh
export GOPATH="$HOME/go"
export PATH="$PATH:$GOPATH/bin"
```

## 包管理

使用go11的modules，请注意go的版本要在1.11或以上。
参考：https://github.com/golang/go/wiki/Modules


#### 激活go modules功能 和 初始化go modules
1. 在$GOPATH/src文件夹之外创建项目。
1. 系统中环境变量GO111MODULE没有设置或者设置为auto。

```bash
go mod init
```
运行go命令时，带着环境变量GO111MODULE=on
```bash
GO111MODULE=on go mod init
```


#### 维护需要的全局运行（global tools）的依赖关系
参考：https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md

* 设置GOBIN, 让go知道去哪里找安装的工具
```bash
export GOBIN=$PWD/bin
export PATH=$GOBIN:$PATH
```

* 追加/bin到gitignore
```bash
$ echo "/bin" >> .gitignore
```

* 在根目录下创建一个tools.go
```go
// +build tools

package tools

import (
	_ "github.com/beinan/gql-server"
)
```

* 安装gql-server工具
```bash
go install github.com/beinan/gql-server
```
* 检查是否真的使用了本module所需版本的工具
在项目目录下
```bash
which gql-server
```
如果一切正常，会看到项目目录下bin的gql-server

* 测试
在main.go中添加：
```go
//go:generate sh -c "gql-server gen model > ./gen/model.go"
//go:generate sh -c "gql-server gen resolver > ./gen/resolver.go"
//go:generate sh -c "gql-server gen gqlresolver > ./gen/gql_resolver.go"
func main(){
  //具体看https://github.com/beinan/gql-server/blob/master/example/main.go
}

```
然后在终端运行
```bash
mkdir gen
go generate
```
#### 日常工作中为项目添加依赖包

* 先import需要的包在.go源代码中
* 运行 go build或者go test

* 如果需要特定版本：
```bash
go get foo@v1.2.3
```
或者直接编辑go.mod文件



