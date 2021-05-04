package usecase

import (
	"context"
	"fmt"
	"net/url"

	"github.com/suzuito/s2-demo-go/entity"
)

func newXMLURLFromArticle(a *entity.ArticleListItem, origin string) *entity.XMLURL {
	return &entity.XMLURL{
		Loc: fmt.Sprintf("%s/article/%s", origin, url.QueryEscape(string(a.ArticleID))),
	}
}

// GenerateBlogSiteMap ...
func GenerateBlogSiteMap(
	ctx context.Context,
	articles *entity.ArticleListItem,
	origin string,
	xmlURLSet *entity.XMLURLSet,
) *entity.XMLURLSet {
	// Articles
	for _, child := range articles.Children {
		xmlURLSet.URLs = append(xmlURLSet.URLs, *newXMLURLFromArticle(&child, origin))
	}

	xmlURLSet.URLs = append(xmlURLSet.URLs, entity.XMLURL{
		Loc: fmt.Sprintf("%s/", origin),
	})

	xmlURLSet.XMLNSXsi = "http://www.w3.org/2001/XMLSchema-instance"
	xmlURLSet.XMLNS = "http://www.sitemaps.org/schemas/sitemap/0.9"
	xmlURLSet.XsiSchemaLocation = "http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd"

	return nil
}
