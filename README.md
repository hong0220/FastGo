## 打包

```
git tag v0.0.1
git push --tags
go get -u github.com/hong0220/FastGo@v0.0.1
```

## 获取

```
go get -u github.com/hong0220/FastGo@v0.0.1
```

## 代码生成

```
1.重新安装 go，再 gf install
2.gf gen dao -c ./hack/config.yarm -path ./internal
3.gf gen service
4.go test gen_test.go
```

## 项目启动

```
program arguments:--gf.gcfg.file=../../manifest/config/config.local.yaml
```


中间件
功能集成
变量枚举
多表联合查询

# TimeWheel
TimeWheel 时间轮是实现延时队列的一种方式，可以优雅执行定时任务。
http://lk668.github.io/2021/04/05/2021-04-05-%E6%89%8B%E6%8A%8A%E6%89%8B%E6%95%99%E4%BD%A0%E5%A6%82%E4%BD%95%E7%94%A8golang%E5%AE%9E%E7%8E%B0%E4%B8%80%E4%B8%AAtimewheel
2.补充资料