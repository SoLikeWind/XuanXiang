package service

import (
	v1 "github.com/SoLikeWind/XuanXiang/api/blog/v1"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlog)

type Blog struct {
	v1.UnimplementedBlogServer //结构体实现了所有 gRPC 方法，但每个方法都返回 Unimplemented 错误

}

func NewBlog() *Blog { return &Blog{} }
