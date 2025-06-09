package service

import (
	"context"

	v1 "github.com/SoLikeWind/XuanXiang/api/blog/v1"
)

// TagService is the server API for Tag service.
func (s *Blog) GetTag(ctx context.Context, req *v1.GetTagReq) (*v1.GetTagReply, error) {
	return &v1.GetTagReply{}, nil
}
