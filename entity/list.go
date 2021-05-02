package entity

import (
	"io"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/xerrors"
)

type ArticleListItem struct {
	Title       string
	Description string
	Children    []ArticleListItem
}

func NewArticleListItemFromHTML(in io.Reader) ([]ArticleListItem, error) {
	ret := []ArticleListItem{}
	var d *goquery.Document
	var err error
	d, err = goquery.NewDocumentFromReader(in)
	if err != nil {
		return nil, xerrors.Errorf("Cannot new goquery : %w", err)
	}
	d.Find("*").Each(func(i int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			switch n.Type {
			case html.ElementNode:
				tag := n.Data
				if tag == "h1" || tag == "h2" || tag == "h3" || tag == "h4" || tag == "h5" {
					if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
						a := ArticleListItem{
							Title:    n.FirstChild.Data,
							Children: []ArticleListItem{},
						}
						ret = append(ret, a)
					}
				}
			}
		}
	})
	return ret, nil
}
