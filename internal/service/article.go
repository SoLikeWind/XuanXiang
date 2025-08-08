package service

import (
	"context"
	"time"

	v1 "github.com/SoLikeWind/XuanXiang/api/blog/v1"
	"github.com/SoLikeWind/XuanXiang/internal/pkg/convert"
	"github.com/SoLikeWind/XuanXiang/internal/pkg/errors"
	"github.com/SoLikeWind/XuanXiang/model/ent"
	entArticle "github.com/SoLikeWind/XuanXiang/model/ent/article"
	"github.com/SoLikeWind/XuanXiang/model/ent/predicate"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *BlogService) ListArticle(ctx context.Context, req *v1.ListArticleReq) (*v1.ListArticleReply, error) {
	articles, total, err := s.article.List(ctx, req.Page, req.PageSize, req.Tag)
	if err != nil {
		return nil, errors.Error(errors.ERROR_LIST_ARTICLE, err)
	}
	return &v1.ListArticleReply{
		Total:    total,
		Articles: convert.EntArticlesToAPI(articles),
	}, nil
}

func (s *BlogService) CreateArticle(ctx context.Context, req *v1.CreateArticleReq) (*v1.CreateArticleReply, error) {
	articleTagsReq := make([]*v1.CreateArticleReq_ArticleTag, 0)
	for _, tag := range req.ArticleTags {
		articleTagsReq := append(articleTagsReq, &v1.CreateArticleReq_ArticleTag{
			Name: tag.Name,
		})
	}

	article, err := s.article.Create(ctx, &ent.Article{ //创建文章实体并返回
		Title:       req.Title,
		Summary:     req.Summary,
		Image:       *req.Image, //指针可为nil(proto:optional)
		ContentMd:   req.ContentMd,
		ContentHTML: convert.MdToHtml(req.ContentMd), // TODO: md to html
		Views:       0,
		// CreatedAt:   timestamppb.Now(),
	})
	if err != nil { //如果创建失败
		return nil, errors.ERROR_CREATE_ARTICLE //返回自定义的服务内部错误，之后不再注释
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
		return nil, errors.ERROR_GET_ARTICLE
	}

	// 异步更新浏览量，不影响文章获取
	go func() {
		updateCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.article.UpdateViewsAsync(updateCtx, req.Id, article.Views)
	}()

	return &v1.GetArticleReply{
		Article: convert.EntArticleToAPI(article), //将ent.Article转换为api.Article
	}, nil
}

func (s *BlogService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleReq) (*emptypb.Empty, error) {
	_, err := s.article.Get(ctx, req.Id) //获取文章
	if err != nil {
		return nil, errors.ERROR_GET_ARTICLE
	}

	// 更新文章基本信息
	if err = s.article.UpdateByMutation(ctx, []predicate.Article{entArticle.IDEQ(req.Id)},
		func(m *ent.ArticleMutation) {
			if req.GetTitle() != "" {
				m.SetTitle(*req.Title)
			}
			if req.Summary != nil {
				m.SetSummary(*req.Summary)
			}
			if req.Image != nil {
				m.SetImage(*req.Image)
			}
			if req.ContentMd != nil {
				m.SetContentMd(*req.ContentMd)
				m.SetContentHTML(convert.MdToHtml(*req.ContentMd))
			}
		}); err != nil {
		return nil, errors.ERROR_UPDATE_ARTICLE
	}

	// 单独处理标签更新
	if req.GetTags() != nil {
		err = s.article.UpdateArticleTags(ctx, req.Id, req.Tags)
		if err != nil {
			return nil, errors.ERROR_UPDATE_ARTICLE
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *BlogService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleReq) (*emptypb.Empty, error) {
	err := s.article.Delete(ctx, req.Id)
	if err != nil {
		return nil, errors.ERROR_DELETE_ARTICLE
	}
	return &emptypb.Empty{}, nil //返回空结构体，proto提供
}
