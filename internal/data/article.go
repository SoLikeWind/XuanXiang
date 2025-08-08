package data

import (
	"context"
	"fmt"

	"github.com/SoLikeWind/XuanXiang/model/ent"
	"github.com/SoLikeWind/XuanXiang/model/ent/article"
	"github.com/SoLikeWind/XuanXiang/model/ent/predicate"
	"github.com/SoLikeWind/XuanXiang/model/ent/tag"
	"github.com/go-kratos/kratos/v2/log"
)

// ArticleRepo 文章数据访问层

// ArticleRepo 文章数据访问层
type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo 创建文章仓库实例
func NewArticleRepo(data *Data, logger log.Logger) *ArticleRepo {
	return &ArticleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/article")),
	}
}

// Create 创建文章
func (r *ArticleRepo) Create(ctx context.Context, article *ent.Article) (*ent.Article, error) {
	return r.data.db.Article.Create().
		SetTitle(article.Title).
		SetSummary(article.Summary).
		SetImage(article.Image).
		SetContentMd(article.ContentMd).
		SetContentHTML(article.ContentHTML).
		SetViews(article.Views).
		Save(ctx)
}

// Get 根据ID获取文章
func (r *ArticleRepo) Get(ctx context.Context, id int64) (*ent.Article, error) {
	return r.data.db.Article.Get(ctx, id)
}

// List 获取文章列表
func (r *ArticleRepo) List(ctx context.Context, page, pageSize int64, tagName string) ([]*ent.Article, int64, error) {
	query := r.data.db.Article.Query()

	// 如果指定了标签，添加标签过滤
	if tagName != "" {
		query = query.Where(article.HasTagsWith(tag.Name(tagName)))
	}

	// 获取总数
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	articles, err := query.
		Limit(int(pageSize)).
		Offset(int((page - 1) * pageSize)).
		Order(ent.Desc(article.FieldCreatedAt)).
		WithTags().
		All(ctx)

	return articles, int64(total), err
}

// UpdateByMutation 使用mutation更新文章
func (r *ArticleRepo) UpdateByMutation(ctx context.Context, predicates []predicate.Article, mutationFn func(*ent.ArticleMutation)) error {
	update := r.data.db.Article.Update().Where(predicates...)
	mutationFn(update.Mutation())
	return update.Exec(ctx)
}

// UpdateViewsAsync 异步更新浏览量
func (r *ArticleRepo) UpdateViewsAsync(ctx context.Context, id int64, currentViews int64) {
	go func() {
		err := r.data.db.Article.UpdateOneID(id).
			SetViews(currentViews + 1).
			Exec(ctx)
		if err != nil {
			r.log.Errorf("failed to update article views: %v", err)
		}
	}()
}

// UpdateArticleTags 更新文章标签
func (r *ArticleRepo) UpdateArticleTags(ctx context.Context, articleID int64, tagNames []string) error {
	// 先删除现有的标签关联
	err := r.data.db.Article.UpdateOneID(articleID).
		ClearTags().
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to clear article tags: %w", err)
	}

	// 为每个标签名称创建或获取标签，然后关联到文章
	for _, tagName := range tagNames {
		// 查找或创建标签
		tagEntity, err := r.data.db.Tag.Query().
			Where(tag.Name(tagName)).
			Only(ctx)

		if err != nil {
			// 标签不存在，创建新标签
			tagEntity, err = r.data.db.Tag.Create().
				SetName(tagName).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to create tag %s: %w", tagName, err)
			}
		}

		// 关联标签到文章
		err = r.data.db.Article.UpdateOneID(articleID).
			AddTags(tagEntity).
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to add tag %s to article: %w", tagName, err)
		}
	}

	return nil
}

// Delete 删除文章
func (r *ArticleRepo) Delete(ctx context.Context, id int64) error {
	return r.data.db.Article.DeleteOneID(id).Exec(ctx)
}
