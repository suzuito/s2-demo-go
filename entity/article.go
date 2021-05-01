package entity

import (
	geojson "github.com/paulmach/orb/geojson"
)

type ArticleID string

type Article struct {
	ID          ArticleID
	Title       string
	Description string
	Blocks      []ArticleBlock
}

type ArticleTOCLevel int

const (
	ArticleTOCLevelH1 ArticleTOCLevel = 1
	ArticleTOCLevelH2 ArticleTOCLevel = 2
	ArticleTOCLevelH3 ArticleTOCLevel = 3
	ArticleTOCLevelH4 ArticleTOCLevel = 4
	ArticleTOCLevelH5 ArticleTOCLevel = 5
)

type ArticleTOC struct {
	Name  string
	Level ArticleTOCLevel
}

type ArticleBlockType string

const (
	ArticleBlockTypeText          ArticleBlockType = "text"
	ArticleBlockTypeSource        ArticleBlockType = "source"
	ArticleBlockTypeSourceAndText ArticleBlockType = "source_and_text"
)

type ArticleBlock struct {
	Type                ArticleBlockType
	Text                string
	Source              string
	SourceResult        string
	SourceResultGeoJSON *geojson.FeatureCollection
}
