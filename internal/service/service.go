package service

import (
	v1 "github.com/SoLikeWind/XuanXiang/api/helloworld/v1"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlog)

type Blog struct {
	v1.UnimplementedGreeterServer
}
