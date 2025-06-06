package data

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/SoLikeWind/XuanXiang/internal/conf"
	"github.com/SoLikeWind/XuanXiang/model/ent"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	db  *ent.Client
	log *log.Helper
	// TODO wrapped database client
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	//为一个数据库驱动程序添加调试和追踪功能，然后使用该驱动程序初始化一个新的ent.Client实例
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		//下面是返回值：驱动
		log.WithContext(ctx).Info(i...) //返回一个helper的拷贝，记录这些参数（即查询信息）
		tracer := otel.Tracer("ent.")   //创建追踪器ent.
		kind := trace.SpanKindServer    //客户端请求操作的跨度的跨度种类,为常数2
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes( //添加与跨生命周期事件相关的属性。
				attribute.String("db.system", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind), //设置跨度类型
		)
		defer span.End() //结束
	})
	client := ent.NewClient(ent.Driver(sqlDrv)) //新建一个ent实例，配置客户端驱动程序。
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	d := &Data{
		db: client,
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		//if err := d.rdb.Close(); err != nil {
		//	log.Error(err)
		//}
	}, nil
}
