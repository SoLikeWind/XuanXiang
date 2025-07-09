package service

import (
	"context"

	v1 "github.com/SoLikeWind/XuanXiang/api/blog/v1"
	"github.com/SoLikeWind/XuanXiang/internal/pkg/convert"
	"github.com/SoLikeWind/XuanXiang/internal/pkg/errors"
	"github.com/SoLikeWind/XuanXiang/model/ent"
)

// TagService is the server API for Tag service.
func (s *BlogService) GetTag(ctx context.Context, req *v1.GetTagReq) (*v1.GetTagReply, error) {
	return &v1.GetTagReply{}, nil
}

func (s *BlogService) ListTag(ctx context.Context, req *v1.ListTagReq) (*v1.ListTagReply, error) {
	tags, total, err := s.tag.List(ctx, req.Page, req.PageSize, req.Name)
	if err != nil {
		return nil, errors.ERROT_LIST_TAG
	}

	return &v1.ListTagReply{
		Total: total,
		Tags:  convert.EntTagsToAPI(tags),
	}, nil
}

func (s *BlogService) CreateTag(ctx context.Context, req *v1.CreateTagReq) (*v1.CreateTagReply, error) {
	tag, err := s.tag.Create(ctx, &ent.Tag{
		Name: req.Name,
	})
	if err != nil {
		return nil, errors.ERROT_CREATE_TAG
	}

	return &v1.CreateTagReply{
		Tag: convert.EntTagToAPI(tag),
	}, nil
}
