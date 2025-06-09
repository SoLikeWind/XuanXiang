package server

import (
	v1 "github.com/SoLikeWind/XuanXiang/api/blog/v1"
	"github.com/SoLikeWind/XuanXiang/internal/conf"
	"github.com/SoLikeWind/XuanXiang/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, blog *service.Blog, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger), //在每次收到 Http 请求时打印详细请求信息
			validate.Validator(),
		),
	}
	if c.Http.Network != "" { //配置文件中配置了网络类型
		opts = append(opts, http.Network(c.Http.Network)) //http类型
	}
	if c.Http.Addr != "" { //配置文件中配置了地址
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil { //配置文件中配置了超时时间
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	httpSrv := http.NewServer(opts...) //创建一个http服务，将之前的opts配置信息传入

	openAPIhander := openapiv2.NewHandler()
	httpSrv.HandlePrefix("/openapi/", openAPIhander) //注册openAPI接口

	v1.RegisterBlogHTTPServer(httpSrv, blog) //注册服务
	return httpSrv
}
