# 安装Golang 

## 下载安装
官网：
https://golang.org/dl/

# 编写Go代码

## GOPATH

把下面的配置加入~/.zshrc
```zsh
export GOPATH="$HOME/go"
export PATH="$PATH:$GOPATH/bin"
```

## 包管理
使用dep， 参考：https://golang.github.io/dep/docs/introduction.html

#### mac安装dep

```zsh
$ brew install dep
$ brew upgrade dep
```

#### 新项目初始化dep

```zsh
$ dep init
```

#### 为项目添加依赖包

```zsh
$ dep ensure -add  github.com/gorilla/mux
```


