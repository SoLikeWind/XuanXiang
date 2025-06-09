package data

import (
	"context"

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

// // ListArticles 列出文章
// func (ar *ArticleRepo) ListArticles(ctx context.Context, after *ent.Article, first *int, before *ent.Article, last *int, orderBy *ent.ArticleOrder) ([]*ent.Article, error) {
// 	ar.log.WithContext(ctx).Info("ListArticles: %v", after)
// 	return ar.data.db.Article.Query().
// 		Paginate(ctx, after, first, before, last,
// 			ent.WithArticleOrder(orderBy),
// 		)
// }

// // CreateArticle 创建文章
// func (ar *ArticleRepo) CreateArticle(ctx context.Context, article *ent.Article) (*ent.Article, error) {
// 	ar.log.WithContext(ctx).Info("CreateArticle: %v", article)
// 	return ar.data.db.Article.Create().SetInput(article).Save(ctx)
// }

// // GetArticle 获取文章
// func (ar *ArticleRepo) GetArticle(ctx context.Context, id int) (*ent.Article, error) {
// 	ar.log.WithContext(ctx).Info("GetArticle: %v", id)
// 	return ar.data.db.Article.Get(ctx, id)
// }

// // UpdateArticle 更新文章
// func (ar *ArticleRepo) UpdateArticle(ctx context.Context, article *ent.Article) (*ent.Article, error) {
// 	ar.log.WithContext(ctx).Info("UpdateArticle: %v", article)
// 	return ar.data.db.Article.UpdateOne(article).Save(ctx)

// }

// DeleteArticle 删除文章
func (ar *ArticleRepo) DeleteArticle(ctx context.Context, id int) error {
	ar.log.WithContext(ctx).Info("DeleteArticle: %v", id)
	return ar.data.db.Article.DeleteOneID(id).Exec(ctx)
}

// CountArticles 计算文章数量
func (ar *ArticleRepo) CountArticles(ctx context.Context) (int, error) {
	ar.log.WithContext(ctx).Info("CountArticles")
	return ar.data.db.Article.Query().Count(ctx)
}
