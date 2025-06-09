package service

import (
	v1 "github.com/SoLikeWind/XuanXiang/api/blog/v1"
	"github.com/SoLikeWind/XuanXiang/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlog)

type BlogService struct {
	v1.UnimplementedBlogServer //结构体实现了所有 gRPC 方法，但每个方法都返回 Unimplemented 错误
	log                        *log.Helper

	article *data.ArticleRepo //文章业务的数据存储
	tag     *data.TagRepo     //标签业务的数据存储
}

// NewBlog 实例化
func NewBlog(
	article *data.ArticleRepo,
	tag *data.TagRepo,
) *BlogService {
	return &BlogService{
		article: article,
		tag:     tag,
	}
}
