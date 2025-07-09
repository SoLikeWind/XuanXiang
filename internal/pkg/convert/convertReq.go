package convert

import (
	"github.com/SoLikeWind/XuanXiang/api/common/model"
	"github.com/SoLikeWind/XuanXiang/model/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func EntArticleToAPI(article *ent.Article) *model.Article {
	if article == nil {
		return nil
	}

	// 将ent.Article赋值给api.Article
	return &model.Article{
		Id:          article.ID,
		Title:       article.Title,
		Summary:     article.Summary,
		Image:       article.Image,
		ContentMd:   article.ContentMd,
		ContentHtml: article.ContentHTML,
		Views:       article.Views,
		CreatedAt:   timestamppb.New(article.CreatedAt),
		UpdatedAt:   timestamppb.New(article.UpdatedAt),
	}
}

// EntArticlesToAPI 将ent.Article切片转换为API响应
func EntArticlesToAPI(articles []*ent.Article) []*model.Article {
	result := make([]*model.Article, 0, len(articles))
	for _, article := range articles {
		result = append(result, EntArticleToAPI(article))
	}
	return result
}

func EntTagToAPI(tag *ent.Tag) *model.Tag {
	if tag == nil {
		return nil
	}

	return &model.Tag{
		Id:   tag.ID,
		Name: tag.Name,
	}
}

func EntTagsToAPI(tags []*ent.Tag) []*model.Tag {
	result := make([]*model.Tag, 0, len(tags))
	for _, tag := range tags {
		result = append(result, EntTagToAPI(tag))
	}
	return result
}
