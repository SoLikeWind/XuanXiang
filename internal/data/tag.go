package data

import (
	"context"

	"github.com/SoLikeWind/XuanXiang/internal/pkg/errors"
	"github.com/SoLikeWind/XuanXiang/model/ent"
	"github.com/SoLikeWind/XuanXiang/model/ent/tag"
	"github.com/go-kratos/kratos/v2/log"
)

type TagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) *TagRepo {
	return &TagRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/tag")),
	}
}

func (t *TagRepo) List(ctx context.Context, page, pageSize int64, name string) ([]*ent.Tag, int64, error) {
	query := t.data.db.Tag.Query()

	if name != "" {
		query = query.Where(tag.NameContainsFold(name)) //+Fold()不区分大小写
	}

	query = query.Order(ent.Desc(tag.FieldCreatedAt))

	total, err := query.Clone().Count(ctx) //统计总数
	if err != nil {
		return nil, 0, errors.ERROR_COUNT_TAG
	}

	tags, err := query.All(ctx)
	if err != nil {
		return nil, 0, errors.ERROR_LIST_TAG
	}

	return tags, int64(total), nil
}

func (t *TagRepo) Create(ctx context.Context, tag *ent.Tag) (*ent.Tag, error) {
	tag, err := t.data.db.Tag.Create().SetName(tag.Name).Save(ctx)
	if err != nil {
		return nil, errors.ERROR_CREATE_TAG
	}
	return tag, nil
}

// GetOrCreateTags 根据标签名称列表查找或创建标签
func (t *TagRepo) GetOrCreateTags(ctx context.Context, tagNames []string) ([]*ent.Tag, error) {
	var tags []*ent.Tag

	for _, tagName := range tagNames {
		// 先尝试查找已存在的标签
		existingTag, err := t.data.db.Tag.Query().Where(tag.NameEQ(tagName)).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				// 标签不存在，创建新标签
				newTag, err := t.data.db.Tag.Create().
					SetName(tagName).
					Save(ctx)
				if err != nil {
					t.log.Errorf("failed creating tag %s: %s", tagName, err)
					return nil, err
				}
				tags = append(tags, newTag)
			} else {
				// 其他错误
				t.log.Errorf("failed querying tag %s: %s", tagName, err)
				return nil, err
			}
		} else {
			// 标签已存在，直接使用
			tags = append(tags, existingTag)
		}
	}

	return tags, nil
}
