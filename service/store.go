package service

import (
	"context"

	"github.com/suzuito/s2-demo-go/entity"
)

type ArticleStore interface {
	GetArticleList(
		ctx context.Context,
		indecies *entity.ArticleListItem,
	) error
	GetArticle(
		ctx context.Context,
		articleID entity.ArticleID,
		article *entity.Article,
	) error

	PutRawFile(
		ctx context.Context,
		bytesSrc []byte,
		contentTypeSrc string,
		pathDst string,
	) error
}
