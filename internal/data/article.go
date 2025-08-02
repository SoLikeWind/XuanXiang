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
	query := ar.data.db.Article.Query() //查询文章

	//“筛选”——只查有某个标签（即有tagname）的文章。
	if tagName != "" { //如果参数有标签名
		query = query.Where(article.HasTagsWith(tag.NameEQ(tagName))) //根据标签名查询文章
	}

	// 预加载标签信息,每篇文章的标签都会被查出来，放到结果里。
	query = query.WithTags() //WithTags() 不会过滤掉没有标签的文章，只是让每篇文章都带上自己的标签信息

	//防止页大小和页码小于0
	if pageSize <= 0 { //页大小
		pageSize = 10
	}
	if page <= 0 { //页码
		page = 1
	}

	offset := (page - 1) * pageSize //偏移量

	total, err := query.Clone().Count(ctx) //统计总数，Clone() 会创建一个新查询，不影响原查询
	// （直接用 query.Count(ctx)，会改变 query 的内部状态，导致后续用同一个 query 查询数据时，结果可能不对。）
	if err != nil {
		ar.log.Errorf("统计文章失败: %s", err)
		return nil, 0, err
	}

	articles, err := query.
		Limit(int(pageSize)).                    //限制查询数量
		Offset(int(offset)).                     //偏移量
		Order(ent.Desc(article.FieldCreatedAt)). //按创建时间降序排序
		All(ctx)

	if err != nil {
		ar.log.Errorf("获取文章失败: %s", err)
		return nil, 0, err
	}

	return articles, int64(total), nil
}

// CreateArticle 创建文章
func (ar *ArticleRepo) Create(ctx context.Context, article *ent.Article) (*ent.Article, error) {
	var tags []*ent.Tag
	for _, tagName := range tags.Name {
		tag, err := ar.data.db.Tag.Query().Where(tag.NameEQ(tagName)).Only(ctx)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag) //将tag添加到article.Tags中
	}

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
	//ar.log.WithContext(ctx).Info("UpdateArticle: %v", article)
	data, err := ar.data.db.Article.UpdateOne(article).Save(ctx)
	if err != nil {
		ar.log.Errorf("更新文章失败: %s", err)
		return nil, err
	}
	return data, nil
}

// DeleteArticle 删除文章
func (ar *ArticleRepo) Delete(ctx context.Context, id int64) error {
	// ar.log.WithContext(ctx).Info("DeleteArticle: %v", id)
	err := ar.data.db.Article.DeleteOneID(id).Exec(ctx)
	if err != nil {
		ar.log.Errorf("删除文章失败: %s", err)
		return err
	}
	return nil
}

// CountArticles 计算文章数量
func (ar *ArticleRepo) Count(ctx context.Context) (int, error) {
	ar.log.WithContext(ctx).Info("CountArticles")
	return ar.data.db.Article.Query().Count(ctx)
}
