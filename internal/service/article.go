package service

import (
	"context"

	v1 "github.com/SoLikeWind/XuanXiang/api/blog/v1"
	"github.com/SoLikeWind/XuanXiang/internal/pkg/convert"
	"github.com/SoLikeWind/XuanXiang/internal/pkg/errors"
	"github.com/SoLikeWind/XuanXiang/model/ent"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *BlogService) ListArticle(ctx context.Context, req *v1.ListArticleReq) (*v1.ListArticleReply, error) {
	articles, total, err := s.article.List(ctx, req.Page, req.PageSize, req.Tag)
	if err != nil {
		return nil, errors.ERROT_LIST_ARTICLE
	}
	return &v1.ListArticleReply{
		Total:    total,
		Articles: convert.EntArticlesToAPI(articles),
	}, nil
}

func (s *BlogService) CreateArticle(ctx context.Context, req *v1.CreateArticleReq) (*v1.CreateArticleReply, error) {
	article, err := s.article.Create(ctx, &ent.Article{ //创建文章实体并返回
		Title:       req.Title,
		Summary:     req.Summary,
		Image:       *req.Image, //指针可为nil(proto:optional)
		ContentMd:   req.ContentMd,
		ContentHTML: convert.MdToHtml(req.ContentMd), // TODO: md to html
		Views:       0,
		// CreatedAt:   timestamppb.Now(),
	})
	if err != nil {
		return nil, errors.ERROT_CREATE_ARTICLE
	}
	return &v1.CreateArticleReply{ //返回创建的文章为api的reply
		Article: convert.EntArticleToAPI(article),
	}, nil
}

func (s *BlogService) GetArticle(ctx context.Context, req *v1.GetArticleReq) (*v1.GetArticleReply, error) {
	// if req.Id < 1 { //参数校验
	// 	return nil, v1.ErrorArticleNotFound("无效的文章id<0")
	// } //不用，因为参数校验在proto里做就行

	tr := otel.Tracer("api")                 //创建一个追踪器，去追踪api
	ctx, span := tr.Start(ctx, "GetArticle") //开始追踪api里的GetArticle,返回一个跨度
	defer span.End()                         //追踪跨度结束

	article, err := s.article.Get(ctx, req.Id)
	if err != nil {
		return nil, errors.ERROT_GET_ARTICLE
	}
	return &v1.GetArticleReply{
		Article: convert.EntArticleToAPI(article), //将ent.Article转换为api.Article
	}, nil
}

func (s *BlogService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleReq) (*emptypb.Empty, error) {
	article, err := s.article.Get(ctx, req.Id)
	if err != nil {
		return nil, errors.ERROT_GET_ARTICLE
	}

	if req.Title != nil {
		article.Title = *req.Title
	}
	if req.Summary != nil {
		article.Summary = *req.Summary
	}
	if req.Image != nil {
		article.Image = *req.Image
	}
	if req.ContentMd != nil {
		article.ContentMd = *req.ContentMd
		article.ContentHTML = convert.MdToHtml(*req.ContentMd)
	}

	_, err = s.article.Update(ctx, article)
	if err != nil {
		return nil, errors.ERROT_UPDATE_ARTICLE
	}
	return &emptypb.Empty{}, nil
}

func (s *BlogService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleReq) (*emptypb.Empty, error) {
	err := s.article.Delete(ctx, req.Id)
	if err != nil {
		return nil, errors.ERROT_DELETE_ARTICLE
	}
	return &emptypb.Empty{}, nil
}
