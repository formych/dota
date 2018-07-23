# dota竞猜
#### Powered By YY @ 2013
#### Repowered @ 2017
### Resources
If you need help getting started with Go itself, I recommend these resources:

> 1.[Go](https://golang.org/)
> 2.[The Go tour](https://tour.golang.org/)
> 3.[How to write Go code](https://golang.org/doc/code.html)
> 4.[Effective Go](https://golang.org/doc/effective_go.html)
>5.[The Go Programming Language](https://github.com/gopl-zh/gopl-zh.github.com)

### Getting Started
```
go get git@github.com:formych/dota.git

cd $GOPATH/src/github.com/formych/dota

go get -u github.com/golang/dep/cmd/dep

dep ensure -v

go run main.go -c release/config/qa.yaml
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
