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
		return nil, 0, errors.ERROT_COUNT_TAG
	}

	tags, err := query.All(ctx)
	if err != nil {
		return nil, 0, errors.ERROT_LIST_TAG
	}

	return tags, int64(total), nil
}

func (t *TagRepo) Create(ctx context.Context, tag *ent.Tag) (*ent.Tag, error) {
	tag, err := t.data.db.Tag.Create().SetName(tag.Name).Save(ctx)
	if err != nil {
		return nil, errors.ERROT_CREATE_TAG
	}
	return tag, nil
}
