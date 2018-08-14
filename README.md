# dota guess

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/formych/dota/master/LICENSE)

#### Powered By YY @ 2013, Repowered @ 2017

### Resources

If you need help getting started with Go itself, I recommend these resources:

> 1. [Go](https://golang.org/)
> 2. [The Go tour](https://tour.golang.org/)
> 3. [How to write Go code](https://golang.org/doc/code.html)
> 4. [Effective Go](https://golang.org/doc/effective_go.html)
> 5. [The Go Programming Language](https://github.com/gopl-zh/gopl-zh.github.com)

### Getting Started

``` sh
go get git@github.com:formych/dota.git

cd $GOPATH/src/github.com/formych/dota

go get -u github.com/golang/dep/cmd/dep

dep ensure -v

go run main.go -c release/config/qa.yaml
```

## build

```bash
sh build.sh
```

### 目录结构

> + api                 对外接口
> - dao                 dao 操作
> * doc                 说明文档
> + config              配置对象
> - model               业务逻辑
> * router              路由对象
> + release/bin         执行文件
> * release/config      配置文件


### enum status
> * -1 删除
> * 0 
> * 1 正常
> *


### todo list

> * 日志切割
> * 规范代码
> * 角色管理
> * 权限控制
> * 数据库读写分离, 事务
> * Redis, 缓存
> * 安全数量控制
> * 数据的强校验


用户登录，目前redis存储唯一key。 jwt包含id和唯一标识
业务估计不大
用户信息存储在Redis?
token里面该包含什么

用redis控制奖金池, 参与人数, 比赛结算同意入库？
 以后加入多种登录方式，注册内容更详细
dao的写法太繁琐，可以优化，后续


均分，按比例分
下拉加载
仅限移动端
输入法遮盖
累了就歇歇，不要放弃
