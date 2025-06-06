package main

import (
	"flag"
	"os"

	"XuanXiang/internal/conf"

	zerolog_kratos "github.com/go-kratos/kratos/contrib/log/zerolog/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/rs/zerolog"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
} //flagconf变量 / 名字 / 如果未指定从conf，用../../configs / 表示用法：-conf config.yaml，如果不加-conf，默认../../configs

// 新建一个app，传入grpc服务和http服务。返回kraots实例
func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),                        //服务的唯一 ID
		kratos.Name(Name),                    //服务名称
		kratos.Version(Version),              //version
		kratos.Metadata(map[string]string{}), //metadata，可以用来存储环境、配置等附加信息
		kratos.Logger(logger),                //设置日志组件，使用传入的 logger
		kratos.Server( //server
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse() //解析命令行参数

	// logger := log.With(zerolog.NewLogger(os.Stdout), //with：日志伴随的其他信息	//将输出记录Stout到新的日志记录器
	// 	"ts", log.DefaultTimestamp, //添加时间戳
	// 	"caller", log.DefaultCaller, //添加调用者信息
	// 	"service.id", id, //添加服务ID
	// 	"service.name", Name, //添加服务名称
	// 	"service.version", Version, //添加服务版本
	// 	"trace.id", tracing.TraceID(), //添加traceID
	// 	"span.id", tracing.SpanID(), //添加spanID
	// )

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	level, err := zerolog.ParseLevel(bc.LogLevel)
	if err != nil {
		panic(err)
	}

	// zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.SetGlobalLevel(level)

	Zerologger := zerolog.New(os.Stdout) //新建一个zerolog的记录器，类型为zerolog.Logger
	logger := zerolog_kratos.NewLogger(&Zerologger)
	logger = log.With(logger, //with：日志伴随的其他信息	//kratos的logger可以接受任意类型的logger
		"ts", log.DefaultTimestamp, //添加时间戳
		"service.id", id, //添加服务ID
		"service.name", Name, //添加服务名称
		"service.version", Version, //添加服务版本
		"trace.id", tracing.TraceID(), //添加traceID
		"span.id", tracing.SpanID(), //添加spanID
	)

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
