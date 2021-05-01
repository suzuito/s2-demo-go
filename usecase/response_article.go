package usecase

import (
	"github.com/paulmach/orb/geojson"
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

type ResponseArticleBlock struct {
	Type                entity.ArticleBlockType    `json:"type"`
	Text                string                     `json:"text"`
	Source              string                     `json:"source"`
	SourceResult        string                     `json:"sourceResult"`
	SourceResultGeoJSON *geojson.FeatureCollection `json:"sourceResultGeoJSON"`
}

func NewResponseArticleBlock(ab *entity.ArticleBlock) *ResponseArticleBlock {
	return &ResponseArticleBlock{
		Type:                ab.Type,
		Text:                ab.Text,
		Source:              ab.Source,
		SourceResult:        ab.SourceResult,
		SourceResultGeoJSON: ab.SourceResultGeoJSON,
	}
}

type ResponseArticleListItem struct {
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Children    []ResponseArticleListItem `json:"children"`
}

func NewResponseArticleListItem(a *entity.ArticleListItem) *ResponseArticleListItem {
	children := []ResponseArticleListItem{}
	for _, b := range a.Children {
		children = append(children, *NewResponseArticleListItem(&b))
	}
	return &ResponseArticleListItem{
		Title:       a.Title,
		Description: a.Description,
		Children:    children,
	}
}
