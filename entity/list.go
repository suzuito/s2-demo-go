package entity

import (
	"io"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/xerrors"
)

var TagNamesForAnchor = []string{"h1", "h2", "h3", "h4", "h5"}

type ArticleListItem struct {
	ArticleID ArticleID
	Anchor    string
	Name      string
	Children  []ArticleListItem
}

func NewArticleListItemFromArticle(a *Article) *ArticleListItem {
	return &ArticleListItem{
		ArticleID: a.ID,
		Anchor:    "",
		Children:  []ArticleListItem{},
		Name:      a.Title,
	}
}

func NewArticleListFromHTML(articleID ArticleID, in io.Reader) ([]ArticleListItem, error) {
	ret := []ArticleListItem{}
	var d *goquery.Document
	var err error
	d, err = goquery.NewDocumentFromReader(in)
	if err != nil {
		return nil, xerrors.Errorf("Cannot new goquery : %w", err)
	}
	d.Find("*").Each(func(i int, s *goquery.Selection) {
		attrID := s.AttrOr("id", "")
		if attrID == "" {
			return
		}
		a := ArticleListItem{
			ArticleID: articleID,
			Anchor:    attrID,
			Children:  []ArticleListItem{},
			Name:      attrID,
		}
		ret = append(ret, a)
	})
	return ret, nil
}
