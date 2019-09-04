# ServerlessPlus

ServerlessPlus 提供了针对目前开放云 serverless 平台的包装，让你可以像本地应用开发一样使用常用 Web 框架开发你的 Serverless 应用，可以让你一套程序同时兼容多种运行环境（普通模式 + Serverless 模式）。

## Status

目前项目可以认为是`pre-alpha`阶段，目前程序不建议在正式业务中使用。目前本框架主要是用于演示思路和整理相关功能列表，如果你有任何想法或建议，可以在 `https://github.com/ipfans/serverlessplus/issues` 中提出。

任何现有 API 均可能在后续版本中进行调整。

### 支持平台产品

- [*] 阿里云函数计算
- [*] 腾讯云无服务器函数
- [ ] Amazon Lambda
- [ ] GCP Cloud Functions

## Demo

示例代码可以在`demo/`目录下找到。

### Beego

```go
func main() {
	b := beego.NewControllerRegister()
	b.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	serverlessplus.Start(b)
}
```

### Echo

```go
func main() {
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})
	serverlessplus.Start(e)
}
```

### Gin

```go
func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
	serverlessplus.Start(r)
}

```
