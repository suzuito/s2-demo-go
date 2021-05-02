package entity

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"

	"golang.org/x/xerrors"
	"gopkg.in/yaml.v2"
)

type ArticleID string

type Article struct {
	ID          ArticleID
	Title       string
	Description string
	Text        string
	Blocks      []ArticleBlock
	PublishedAt int64
	Draft       bool
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

type ArticleBlockID string

type ArticleBlock struct {
	ID                      ArticleBlockID
	Type                    ArticleBlockType
	PathText                string
	PathSource              string
	PathSourceResult        string
	PathSourceResultGeoJSON string
}

var ErrMetaBlockNotFound = fmt.Errorf("Article meta block not found")

// NewArticleFromRawContent ...
func NewArticleFromRawContent(r io.Reader) (*Article, []byte, error) {
	s := bufio.NewScanner(r)
	isMetaBlock := false
	isMetaBlockDone := false
	metaBlock := ""
	notMetaBlock := ""
	for s.Scan() {
		l := s.Text()
		if strings.HasPrefix(l, "---") && !isMetaBlockDone {
			if !isMetaBlock {
				isMetaBlock = true
				continue
			}
			isMetaBlock = false
			isMetaBlockDone = true
			continue
		}
		if isMetaBlock {
			metaBlock += l + "\n"
		} else {
			notMetaBlock += l + "\n"
		}
	}
	if !isMetaBlockDone {
		return nil, nil, xerrors.Errorf("Meta data is not found : %w", ErrMetaBlockNotFound)
	}
	embedMeta := struct {
		ID          string   `yaml:"id"`
		Title       string   `yaml:"title"`
		Tags        []string `yaml:"tags"`
		Description string   `yaml:"description"`
		Date        string   `yaml:"date"`
		Draft       bool     `yaml:"draft"`
	}{}
	if err := yaml.Unmarshal([]byte(metaBlock), &embedMeta); err != nil {
		return nil, nil, xerrors.Errorf("Cannot parse yaml block '%s' : %w", metaBlock, err)
	}
	date, err := time.Parse("2006-01-02", embedMeta.Date)
	if err != nil {
		return nil, nil, xerrors.Errorf("Cannot parse date '%s' : %w", embedMeta.Date, err)
	}
	article := Article{
		ID:          ArticleID(embedMeta.ID),
		Title:       embedMeta.Title,
		Description: embedMeta.Description,
		PublishedAt: date.Unix(),
		Draft:       embedMeta.Draft,
	}
	return &article, []byte(notMetaBlock), nil
}

// NewArticleBlockFromRawContent ...
func NewArticleBlockFromRawContent(r io.Reader) (*ArticleBlock, []byte, error) {
	s := bufio.NewScanner(r)
	isMetaBlock := false
	isMetaBlockDone := false
	metaBlock := ""
	notMetaBlock := ""
	for s.Scan() {
		l := s.Text()
		if strings.HasPrefix(l, "---") && !isMetaBlockDone {
			if !isMetaBlock {
				isMetaBlock = true
				continue
			}
			isMetaBlock = false
			isMetaBlockDone = true
			continue
		}
		if isMetaBlock {
			metaBlock += l + "\n"
		} else {
			notMetaBlock += l + "\n"
		}
	}
	if !isMetaBlockDone {
		return nil, nil, xerrors.Errorf("Meta data is not found : %w", ErrMetaBlockNotFound)
	}
	embedMeta := struct {
		ID string `yaml:"id"`
	}{}
	if err := yaml.Unmarshal([]byte(metaBlock), &embedMeta); err != nil {
		return nil, nil, xerrors.Errorf("Cannot parse yaml block '%s' : %w", metaBlock, err)
	}
	article := ArticleBlock{
		ID: ArticleBlockID(embedMeta.ID),
	}
	return &article, []byte(notMetaBlock), nil
}
