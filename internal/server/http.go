package server

import (
	"fmt"
	net_http "net/http"

	v1 "github.com/SoLikeWind/XuanXiang/api/blog/v1"
	"github.com/SoLikeWind/XuanXiang/internal/conf"
	"github.com/SoLikeWind/XuanXiang/internal/service"
	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,
	logger log.Logger,
	blog *service.BlogService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger), //在每次 Http 请求后调用服务,打印对应crud服务的 相关信息
			validate.ProtoValidate(),
		),
		http.Filter( //过滤器
			handlers.CORS( //设置允许的跨域请求方法
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedHeaders([]string{
					"Accept", "Accept-Language", "Content-Language", "Origin",
					"Access-Control-Request-Method",
					"Access-Control-Request-Headers",
					"Content-Type",
					"Authorization",
					"X-Request-Id",
					"User-Agent",
				}),
				handlers.ExposedHeaders([]string{
					"Content-Type",
					"Origin",
					"b3",
					"User-Agent",
					"X-Requested-With",
					"X-Request-Id",
					"X-Meta",
					"X-RateLimit-Limit",
					"X-RateLimit-Remaining",
					"X-RateLimit-Reset",
				}),
				handlers.AllowedMethods([]string{
					net_http.MethodGet,
					net_http.MethodPost,
					net_http.MethodHead,
					net_http.MethodConnect,
					net_http.MethodPut,
					net_http.MethodPatch,
					net_http.MethodDelete,
					net_http.MethodOptions,
				}),
			),
		),
	}
	if c.Http.Network != "" { //配置文件中配置了网络类型
		opts = append(opts, http.Network(c.Http.Network)) //http类型
		fmt.Println("http network:", c.Http.Network)
	}
	if c.Http.Addr != "" { //配置文件中配置了地址
		opts = append(opts, http.Address(c.Http.Addr))
		fmt.Println("http addr:", c.Http.Addr)
	}
	if c.Http.Timeout != nil { //配置文件中配置了超时时间
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
		fmt.Println("http timeout:", c.Http.Timeout.AsDuration())
	}
	httpSrv := http.NewServer(opts...) //创建一个http服务，将之前的opts配置信息传入

	openAPIhander := openapiv2.NewHandler()    //创建一个openAPI处理器
	httpSrv.HandlePrefix("/q/", openAPIhander) //注册openAPI接口

	v1.RegisterBlogHTTPServer(httpSrv, blog) //注册服务
	return httpSrv
}
