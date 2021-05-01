package usecase

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday/v2"
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

var regexpFileExtReplacerMarkdown = regexp.MustCompile(".md$")

func ConvertFileMarkdownToHTML(
	srcFilePath string,
) error {
	src, err := os.Open(srcFilePath)
	if err != nil {
		return xerrors.Errorf("Cannot open src '%s' : %w", srcFilePath, err)
	}
	defer src.Close()
	dstFilePath := regexpFileExtReplacerMarkdown.ReplaceAllString(srcFilePath, ".html")
	dst, err := os.Open(dstFilePath)
	if err != nil {
		return xerrors.Errorf("Cannot open src '%s' : %w", dstFilePath, err)
	}
	defer dst.Close()
	return ConvertMarkdownToHTML(src, dst)
}

func ConvertMarkdownToHTML(
	srcMarkdown io.Reader,
	dstHTML io.Writer,
) error {
	bytesMarkdown, err := ioutil.ReadAll(srcMarkdown)
	if err != nil {
		return xerrors.Errorf("Cannot read src markdown : %w", err)
	}
	bytesHTML1 := blackfriday.Run(
		bytesMarkdown,
		blackfriday.WithRenderer(
			blackfriday.NewHTMLRenderer(
				blackfriday.HTMLRendererParameters{
					Flags: blackfriday.TOC | blackfriday.HrefTargetBlank,
				},
			),
		),
	)
	bytesHTML2, err := convertAfterConvert(bytesHTML1)
	if err != nil {
		return xerrors.Errorf("Cannot convertAfterConvert : %w", err)
	}
	fmt.Fprintf(dstHTML, bytesHTML2)
	return nil
}

func convertAfterConvert(body []byte) (string, error) {
	d, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return "", xerrors.Errorf("Cannot new goquery : %w", err)
	}
	d.Find("pre").Each(func(i int, s *goquery.Selection) {
		s.SetAttr("class", "code-block")
		s.SetAttr("style", "width: 100%; overflow: scroll;")
	})
	returned, err := d.Html()
	if err != nil {
		return "", xerrors.Errorf("Cannot create goquery html : %w", err)
	}
	returned = strings.Replace(returned, "<html><head></head><body>", "", 1)
	returned = strings.Replace(returned, "</body></html>", "", 1)
	return returned, nil
}
