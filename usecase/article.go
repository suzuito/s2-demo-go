package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

func BuildArticleBlock(
	bodyMarkdown []byte,
	bodyHTML *[]byte,
	block *entity.ArticleBlock,
) error {
	tmp, convertedMarkdown, err := entity.NewArticleBlockFromRawContent(bytes.NewReader(bodyMarkdown))
	if err != nil {
		return xerrors.Errorf("Cannot NewArticleBlockFromRawContent : %w", err)
	}
	if err := ConvertMarkdownToHTML(convertedMarkdown, bodyHTML); err != nil {
		return xerrors.Errorf("Cannot convert : %w", err)
	}
	*block = *tmp
	return nil
}

func BuildArticle(
	bodyMarkdown []byte,
	article *entity.Article,
	bodyHTML *[]byte,
) error {
	tmp, convertedMarkdown, err := entity.NewArticleFromRawContent(bytes.NewReader(bodyMarkdown))
	if err != nil {
		return xerrors.Errorf("Cannot new article : %w", err)
	}
	if err = ConvertMarkdownToHTML(convertedMarkdown, bodyHTML); err != nil {
		return xerrors.Errorf("Cannot build md to html : %w", err)
	}
	*article = *tmp
	return nil
}

func BuildArticleToLocal(
	dirPath string,
) error {
	var err error
	outputBytes := map[string][]byte{}
	outputObjects := map[string]interface{}{}
	articleMarkdownBytes := []byte{}
	articleMarkdownBytes, err = ioutil.ReadFile(filepath.Join(dirPath, "article.md"))
	if err != nil {
		return xerrors.Errorf("Cannot ReadFile : %w", err)
	}
	articleHTMLBytes := []byte{}
	article := entity.Article{}
	if err := BuildArticle(articleMarkdownBytes, &article, &articleHTMLBytes); err != nil {
		return xerrors.Errorf("Cannot build article : %w", err)
	}
	outputBytes[filepath.Join(dirPath, "article.html")] = articleHTMLBytes
	if err != nil {
		return xerrors.Errorf("Cannot new article list item : %w", err)
	}
	if err := filepath.Walk(dirPath, func(dirPathArticleBlock string, info1 os.FileInfo, _ error) error {
		if !info1.IsDir() {
			return nil
		}
		left, _ := filepath.Abs(dirPathArticleBlock)
		right, _ := filepath.Abs(dirPath)
		if left == right {
			return nil
		}
		block := entity.ArticleBlock{}
		blockMarkdownBytes := []byte{}
		blockMarkdownBytes, err = ioutil.ReadFile(filepath.Join(dirPathArticleBlock, "article.md"))
		if err != nil {
			return xerrors.Errorf("Cannot ReadFile : %w", err)
		}
		blockHTMLBytes := []byte{}
		if err := BuildArticleBlock(blockMarkdownBytes, &blockHTMLBytes, &block); err != nil {
			return xerrors.Errorf("Cannot build article block : %w", err)
		}
		outputBytes[filepath.Join(dirPathArticleBlock, "article.html")] = blockHTMLBytes
		article.Blocks = append(article.Blocks, block)
		return nil
	}); err != nil {
		return xerrors.Errorf("Cannot walk : %w", err)
	}
	outputObjects[filepath.Join(dirPath, "article.json")] = NewResponseArticle(&article)
	if err := writeFileBytes(outputBytes); err != nil {
		return xerrors.Errorf("Cannot writeFileBytes : %w", err)
	}
	if err := writeFileObjects(outputObjects); err != nil {
		return xerrors.Errorf("Cannot writeFileObjects : %w", err)
	}
	return nil
}

func writeFileBytes(files map[string][]byte) error {
	for name, f := range files {
		if err := ioutil.WriteFile(name, f, 0644); err != nil {
			return xerrors.Errorf("Cannot WriteFile : %w", err)
		}
	}
	return nil
}

func writeFileObjects(files map[string]interface{}) error {
	for name, f := range files {
		b, err := json.Marshal(f)
		if err != nil {
			return xerrors.Errorf("Cannot marshal : %w", err)
		}
		if err := ioutil.WriteFile(name, b, 0644); err != nil {
			return xerrors.Errorf("Cannot WriteFile : %w", err)
		}
	}
	return nil
}

func BuildIndexToLocal(
	dirPath string,
) error {
	if err := filepath.Walk(dirPath, func(dirPathArticleBlock string, info1 os.FileInfo, _ error) error {
		fmt.Println(dirPath)
		return nil
	}); err != nil {
		return xerrors.Errorf("Cannot walk : %w", err)
	}
	return nil
}
