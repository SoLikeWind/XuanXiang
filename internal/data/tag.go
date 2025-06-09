package data

import "github.com/go-kratos/kratos/v2/log"

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
