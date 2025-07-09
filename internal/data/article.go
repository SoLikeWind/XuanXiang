package data

import (
	"context"

	"github.com/SoLikeWind/XuanXiang/model/ent"
	"github.com/SoLikeWind/XuanXiang/model/ent/article"
	"github.com/SoLikeWind/XuanXiang/model/ent/tag"
	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo 实例化
func NewArticleRepo(data *Data, logger log.Logger) *ArticleRepo {
	return &ArticleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/article")),
	}
}

// ListArticles 列出文章
func (ar *ArticleRepo) List(ctx context.Context, page, pageSize int64, tagName string) ([]*ent.Article, int64, error) {
	query := ar.data.db.Article.Query()

	if tagName != "" {
		query = query.Where(article.HasTagsWith(tag.NameEQ(tagName)))
	}

	// 预加载标签信息
	query = query.WithTags()

	if pageSize <= 0 { //页大小
		pageSize = 10
	}
	if page <= 0 { //页码
		page = 1
	}

	offset := (page - 1) * pageSize //偏移量

	total, err := query.Clone().Count(ctx)
	if err != nil {
		ar.log.Errorf("failed counting articles: %s", err)
		return nil, 0, err
	}

	articles, err := query.
		Limit(int(pageSize)).                    //限制查询数量
		Offset(int(offset)).                     //偏移量
		Order(ent.Desc(article.FieldCreatedAt)). //按创建时间降序排序
		All(ctx)

	if err != nil {
		ar.log.Errorf("failed listing articles: %s", err)
		return nil, 0, err
	}

	return articles, int64(total), nil
}

// CreateArticle 创建文章
func (ar *ArticleRepo) Create(ctx context.Context, article *ent.Article) (*ent.Article, error) {
	data, err := ar.data.db.Article.Create().
		SetTitle(article.Title).
		SetSummary(article.Summary).
		SetImage(article.Image).
		SetContentMd(article.ContentMd).
		SetContentHTML(article.ContentHTML).
		SetViews(article.Views).
		// SetCreatedAt(article.CreatedAt).
		// SetUpdatedAt(article.UpdatedAt).
		Save(ctx)
	if err != nil {
		if !ent.IsConstraintError(err) {
			ar.log.Errorf("failed creating article: %s", err)
		}
		return nil, err
	}
	return data, nil
}

// GetArticle 获取文章
func (ar *ArticleRepo) Get(ctx context.Context, id int64) (*ent.Article, error) {
	data, err := ar.data.db.Article.Query().Where(article.IDEQ(id)).Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			ar.log.Errorf("failed getting article: %s", err)
			return nil, err
		}
	}
	return data, nil
}

// UpdateArticle 更新文章
func (ar *ArticleRepo) Update(ctx context.Context, article *ent.Article) (*ent.Article, error) {
	ar.log.WithContext(ctx).Info("UpdateArticle: %v", article)
	return ar.data.db.Article.UpdateOne(article).Save(ctx)
}

// DeleteArticle 删除文章
func (ar *ArticleRepo) Delete(ctx context.Context, id int64) error {
	ar.log.WithContext(ctx).Info("DeleteArticle: %v", id)
	return ar.data.db.Article.DeleteOneID(id).Exec(ctx)
}

// CountArticles 计算文章数量
func (ar *ArticleRepo) Count(ctx context.Context) (int, error) {
	ar.log.WithContext(ctx).Info("CountArticles")
	return ar.data.db.Article.Query().Count(ctx)
}
