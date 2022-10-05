package main

//
//import (
//	"context"
//	"github.com/gogf/gf/v2/frame/g"
//	"github.com/gogf/gf/v2/net/ghttp"
//	"github.com/gogf/gf/v2/os/gcfg"
//	"github.com/gogf/gf/v2/os/gcmd"
//	"github.com/gogf/gf/v2/os/gctx"
//	_ "github.com/hong0220/fastgo/internal/logic"
//	"github.com/hong0220/fastgo/pkg/common/base"
//)
//
//var ctx = gctx.New()
//
//func main() {
//	// 服务名
//	serviceName, _ := g.Cfg().Get(ctx, "serviceName")
//	switch serviceName.Val() {
//	case "api":
//		StartAlarmApi()
//	default:
//		ApiConfigSetUp()
//		StartAlarmApi()
//	}
//	return
//}
//
//func StartAlarmApi() {
//	// 注册日志handler
//	//g.Log().SetHandlers(middleware.LoggingStandHandler)
//
//	// 启动服务
//	command := gcmd.Command{
//		Name:  "main",
//		Usage: "main",
//		Brief: "start http server",
//		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
//			s := g.Server()
//			s.Group("/", func(group *ghttp.RouterGroup) {
//				// 设置跨域问题
//				group.Middleware(MiddlewareCORS)
//				// 响应数据处理
//				group.Middleware(ghttp.MiddlewareHandlerResponse)
//				// 打印 http 请求信息
//				group.Middleware(MiddlewareRequestPrint)
//
//				group.Group("/", func(group *ghttp.RouterGroup) {
//					group.Bind(
//					//controller.AlarmCallback,
//					)
//				})
//
//				group.Group("/", func(group *ghttp.RouterGroup) {
//					// 获取用户信息
//					//group.Middleware(middleware.MiddlewareHandlerAuthRequest)
//					group.Bind(
//					//controller.Alarm,
//					)
//				})
//			})
//
//			// todo
//			//s.Logger().SetHandlers(middleware.LoggingStandHandler)
//			s.GetOpenApi().Info.Title = "fastgo-API文档"
//			s.Run()
//			return nil
//		},
//	}
//	command.Run(ctx)
//}
//
//// 规范路由自动生成  http://gf.l11.cn/  （可参考）
//
//func ApiConfigSetUp() {
//	path := base.GetProjectPath()
//	path = path + "/manifest/config/fastgo-local.yaml"
//	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(path)
//}
//
//func MiddlewareCORS(r *ghttp.Request) {
//	corsOptions := r.Response.DefaultCORSOptions()
//	// 重写header
//	corsOptions.AllowHeaders = "Origin,Content-Type,Accept,Accept-Encoding,Accept-Language,Connection,Host,User-Agent,Cookie,Referer,Authorization,X-Auth-Token,X-Requested-With,X-TENANT-ID"
//	r.Response.CORS(corsOptions)
//	r.Middleware.Next()
//}
//
//func MiddlewareRequestPrint(r *ghttp.Request) {
//	g.Log().Info(r.GetCtx(), "url="+r.Request.RequestURI+
//		",method="+r.Method+
//		",request="+r.GetBodyString()+
//		",clientIp="+r.GetClientIp())
//	r.Middleware.Next()
//}
