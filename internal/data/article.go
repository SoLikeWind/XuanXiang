package data

import (
	"context"

	"github.com/SoLikeWind/XuanXiang/model/ent"
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
// func (ar *ArticleRepo) ListArticles(ctx context.Context, page, pageSize int64, tag string) ([]*ent.Article, error) {
// 	query := ar.data.db.Article.Query()

// 	if tag != "" {
// 		query = query.Where(article.HasTagsWith(tag.NameEQ(tag)))
// 	}

// 	if pageSize <= 0 {
// 		pageSize = 10
// 	}
// 	if page <= 0 {
// 		page = 1
// 	}

// 	offset := (page - 1) * pageSize

// 	articles, err := query.
// 		Limit(int(pageSize)).
// 		Offset(int(offset)).
// 		Order(ent.Desc(article.FieldCreatedAt)).
// 		All(ctx)

// 	if err != nil {
// 		ar.log.Errorf("failed listing articles: %s", err)
// 		return nil, err
// 	}

// 	return articles, nil
// }

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
	ar.log.WithContext(ctx).Info("GetArticle: %v", id)
	return ar.data.db.Article.Get(ctx, id)
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
