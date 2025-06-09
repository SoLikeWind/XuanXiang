package service

import (
	"context"

	v1 "github.com/SoLikeWind/XuanXiang/api/blog/v1"
)

func (s *BlogService) CreateArticle(context *context.Context, req *v1.CreateArticleReq) (*v1.CreateArticleReply, error) {
	article := &Article{
		Title: req.Title,
	}
	return nil, nil
}

func (s *BlogService) ListArticles(context *context.Context, req *v1.ListArticleReq) (*v1.ListArticleReply, error) {

	return nil, nil
}

func (s *BlogService) UpdateArticle(context *context.Context, req *v1.UpdateArticleReq) (*v1.UpdateArticleReply, error) {
	{

		return nil, nil
	}

}

func (s *BlogService) DeleteArticle(context *context.Context, req *v1.DeleteArticleReq) error {

	return nil
}
