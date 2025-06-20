package data

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
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
var ProviderSet = wire.NewSet(NewData, NewArticleRepo, NewTagRepo)

// Data .
type Data struct {
	db  *ent.Client //包含ent.Client实例，schema信息，数据库连接信息
	log *log.Helper
	// TODO wrapped database client
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "data")) //增加一个日志记录器
	drv, err := sql.Open(                                    //sql.Open()函数用于打开数据库连接，返回一个*sql.DB实例
		conf.Database.Driver,
		conf.Database.Source,
	)
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) { //dialect.DebugWithContext()函数用于设置数据库连接的调试模式，
		log.WithContext(ctx).Debug(i...) //并使用log.WithContext()函数将日志记录器添加到上下文中
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
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
		db:  client,
		log: log,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
