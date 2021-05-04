package usecase

import (
	"github.com/suzuito/s2-demo-go/entity"
)

type ResponseArticle struct {
	ID          entity.ArticleID       `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Blocks      []ResponseArticleBlock `json:"blocks"`
}

func NewResponseArticle(a *entity.Article) *ResponseArticle {
	blocks := []ResponseArticleBlock{}
	for _, b := range a.Blocks {
		blocks = append(blocks, *NewResponseArticleBlock(&b))
	}
	return &ResponseArticle{
		ID:          a.ID,
		Title:       a.Title,
		Description: a.Description,
		Blocks:      blocks,
	}
}

func (r *ResponseArticle) Entity() *entity.Article {
	blocks := []entity.ArticleBlock{}
	for _, a := range r.Blocks {
		blocks = append(blocks, *a.Entity())
	}
	return &entity.Article{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		Blocks:      blocks,
	}
}

type ResponseArticleBlock struct {
	ID                  entity.ArticleBlockID   `json:"id"`
	Type                entity.ArticleBlockType `json:"type"`
	Text                string                  `json:"text"`
	Source              string                  `json:"source"`
	SourceResult        string                  `json:"sourceResult"`
	SourceResultGeoJSON string                  `json:"sourceResultGeoJSON"`
}

func NewResponseArticleBlock(ab *entity.ArticleBlock) *ResponseArticleBlock {
	return &ResponseArticleBlock{
		ID:                  ab.ID,
		Type:                ab.Type,
		Text:                ab.PathText,
		Source:              ab.PathSource,
		SourceResult:        ab.PathSourceResult,
		SourceResultGeoJSON: ab.PathSourceResultGeoJSON,
	}
}

func (r *ResponseArticleBlock) Entity() *entity.ArticleBlock {
	return &entity.ArticleBlock{
		ID:                      r.ID,
		Type:                    r.Type,
		PathText:                r.Text,
		PathSource:              r.Source,
		PathSourceResult:        r.SourceResult,
		PathSourceResultGeoJSON: r.SourceResultGeoJSON,
	}
}

type ResponseArticleListItem struct {
	ArticleID entity.ArticleID          `json:"articleId"`
	Anchor    string                    `json:"anchor"`
	Name      string                    `json:"name"`
	Children  []ResponseArticleListItem `json:"children"`
}

func NewResponseArticleListItem(a *entity.ArticleListItem) *ResponseArticleListItem {
	children := []ResponseArticleListItem{}
	for _, b := range a.Children {
		children = append(children, *NewResponseArticleListItem(&b))
	}
	return &ResponseArticleListItem{
		ArticleID: a.ArticleID,
		Anchor:    a.Anchor,
		Name:      a.Name,
		Children:  children,
	}
}
