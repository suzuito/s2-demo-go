package usecase

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday/v2"
	"golang.org/x/xerrors"
)

var regexpFileExtReplacerMarkdown = regexp.MustCompile(".md$")

func replaceExtFileMarkdownToHTML(filePath string) string {
	return regexpFileExtReplacerMarkdown.ReplaceAllString(filePath, ".html")
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

func ConvertMarkdownToHTML(
	srcMarkdown []byte,
	dstHTML *[]byte,
) error {
	bytesHTML1 := blackfriday.Run(
		srcMarkdown,
		blackfriday.WithRenderer(
			blackfriday.NewHTMLRenderer(
				blackfriday.HTMLRendererParameters{
					Flags: blackfriday.HrefTargetBlank,
				},
			),
		),
	)
	bytesHTML2, err := convertAfterConvert(bytesHTML1)
	if err != nil {
		return xerrors.Errorf("Cannot convertAfterConvert : %w", err)
	}
	*dstHTML = []byte(bytesHTML2)
	return nil
}
