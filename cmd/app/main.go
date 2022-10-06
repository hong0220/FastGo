package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/hong0220/FastGo/internal/controller"
	"github.com/hong0220/FastGo/pkg/middleware/metric"
	"time"
)

var ctx = gctx.New()

func main() {
	// 注册日志handler
	g.Log().SetHandlers(LoggingStandHandler)

	//// 初始化自定义校验规则
	//middleware.InitRegisterValidRule(ctx)

	// 初始化 metric
	go metric.Init()

	// 启动服务
	command := gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				// 设置跨域问题
				group.Middleware(MiddlewareCORS)
				// 打印 http 请求信息
				group.Middleware(MiddlewareRequestPrint)
				// 响应数据处理
				group.Middleware(ghttp.MiddlewareHandlerResponse)

				// 不需要校验用户信息
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Bind(
					//controller.AlarmCallback,
					)
				})

				// 需要校验用户信息
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 获取用户信息
					//group.Middleware(middleware.MiddlewareHandlerAuthRequest)
					group.Bind(
						controller.CSysDict,
					)
				})
			})

			//middleware.XXLJob.XxlJobMux(s)

			s.Logger().SetHandlers(LoggingStandHandler)
			s.GetOpenApi().Info.Title = "FastGo 1.0-API文档"
			s.Run()
			return nil
		},
	}
	command.Run(ctx)
}

// [格式版本] [时间] [日志级别] [协程] logger [traceId] message
var LoggingStandHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
	// 按照规范格式定义输出内容
	in.Buffer.WriteString("[v1] ")
	if "" == in.TimeFormat {
		in.Buffer.WriteString("[" + time.Now().Format("2008-08-08 08:08:08.888") + "] ")
	} else {
		in.Buffer.WriteString("[" + in.TimeFormat + "] ")
	}
	if "" == in.LevelFormat {
		in.Buffer.WriteString("[INFO] ")
	} else {
		in.Buffer.WriteString(in.LevelFormat + " ")
	}
	in.Buffer.WriteString("[] ")
	in.Buffer.WriteString("logger ")
	in.Buffer.WriteString("[" + gstr.Trim(in.CtxStr, "{}") + "] ")
	in.Buffer.WriteString(gstr.Trim(in.Content))
	in.Buffer.WriteString("\n")
	in.Next(ctx)
}

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	// 重写header
	corsOptions.AllowHeaders = "Origin,Content-Type,Accept,Accept-Encoding,Accept-Language,Connection,Host,User-Agent,Cookie,Referer,Authorization,X-Auth-Token,X-Requested-With,X-TENANT-ID"
	r.Response.CORS(corsOptions)

	r.Middleware.Next()
}

func MiddlewareRequestPrint(r *ghttp.Request) {
	g.Log().Info(r.GetCtx(), "url="+r.Request.RequestURI+
		",method="+r.Method+
		",request="+r.GetBodyString()+
		",clientIp="+r.GetClientIp())

	r.Middleware.Next()
}
