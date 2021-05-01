package usecase

import (
	"context"

	"github.com/suzuito/s2-demo-go/entity"
	"github.com/suzuito/s2-demo-go/service"
	"golang.org/x/xerrors"
)

func GetArticle(
	ctx context.Context,
	articleStore service.ArticleStore,
	articleID entity.ArticleID,
	article *entity.Article,
) error {
	if err := articleStore.GetArticle(ctx, articleID, article); err != nil {
		return xerrors.Errorf("Cannot get article '%s' : %w", articleID, err)
	}
	return nil
}

func GetArticleList(
	ctx context.Context,
	articleStore service.ArticleStore,
	articleItems *entity.ArticleListItem,
) error {
	if err := articleStore.GetArticleList(ctx, articleItems); err != nil {
		return xerrors.Errorf("Cannot get article list : %w", err)
	}
	return nil
}
