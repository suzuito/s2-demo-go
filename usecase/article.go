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
	root *entity.ArticleListItem,
) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return xerrors.Errorf("Cannot open dir '%s' : %w", dirPath, err)
	}
	for _, file1 := range files {
		if !file1.IsDir() {
			continue
		}
		dirPathArticle := filepath.Join(dirPath, file1.Name())
		filesArticle, err := ioutil.ReadDir(dirPathArticle)
		if err != nil {
			return xerrors.Errorf("Cannot open dir '%s' : %w", dirPathArticle, err)
		}
		var article *entity.Article
		var articleHTMLBytes []byte
		blockArticleHTMLBytesList := [][]byte{}
		for _, file2 := range filesArticle {
			filePath := filepath.Join(dirPathArticle, file2.Name())
			if file2.IsDir() {
				blockArticleHTMLBytes, err := readBlockArticleBytes(filePath)
				if err != nil {
					return xerrors.Errorf("Cannot read : %w", err)
				}
				if blockArticleHTMLBytes != nil {
					blockArticleHTMLBytesList = append(blockArticleHTMLBytesList, blockArticleHTMLBytes)
				}
				continue
			}
			if file2.Name() != "article.json" && file2.Name() != "article.html" {
				continue
			}
			fileBytes, err := ioutil.ReadFile(filePath)
			if err != nil {
				return xerrors.Errorf("Cannot open file '%s' : %w", filePath, err)
			}
			switch file2.Name() {
			case "article.json":
				if err := newArticleFromFile(filePath, article); err != nil {
					return xerrors.Errorf("Cannot newArticleFrom file '%s' : %w", filePath, err)
				}
			case "article.html":
				articleHTMLBytes = fileBytes
			default:
				return xerrors.Errorf("Unsupported file '%s'", filePath)
			}
		}
		if article == nil {
			return xerrors.Errorf("Article is not found on dir '%s'", dirPathArticle)
		}
		articleListItem := entity.NewArticleListItemFromArticle(article)
		if articleHTMLBytes != nil {
			children, err := entity.NewArticleListFromHTML(articleListItem.ArticleID, bytes.NewReader(articleHTMLBytes))
			if err != nil {
				return xerrors.Errorf("Cannot new article list from html : %w", err)
			}
			articleListItem.Children = children
		}
		for _, blockArticleHTMLBytes := range blockArticleHTMLBytesList {
			children, err := entity.NewArticleListFromHTML(articleListItem.ArticleID, bytes.NewReader(blockArticleHTMLBytes))
			if err != nil {
				return xerrors.Errorf("Cannot new article list from html : %w", err)
			}
			articleListItem.Children = append(articleListItem.Children, children...)
		}
		root.Children = append(root.Children, *articleListItem)
	}
	return nil
}

func readBlockArticleBytes(
	dirPathBlock string,
) ([]byte, error) {
	filePathBlockArticle := filepath.Join(dirPathBlock, "article.html")
	if _, err := os.Stat(filePathBlockArticle); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, xerrors.Errorf("Cannot stat '%s' : %w", filePathBlockArticle, err)
	}
	bytesBlockArticle, err := ioutil.ReadFile(filePathBlockArticle)
	if err != nil {
		return nil, xerrors.Errorf("Cannot open file '%s' : %w", filePathBlockArticle, err)
	}
	return bytesBlockArticle, nil
}

func newArticleFromFile(filePath string, article *entity.Article) error {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return xerrors.Errorf("Cannot read file '%s' : %w", filePath, err)
	}
	rArticle := ResponseArticle{}
	if err := json.Unmarshal(fileBytes, &rArticle); err != nil {
		return xerrors.Errorf("Cannot unmarshal file '%s' : %w", filePath, err)
	}
	*article = *rArticle.Entity()
	return nil
}

func UploadArticles(
	ctx context.Context,
	articleStore service.ArticleStore,
	dirPath string,
) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return xerrors.Errorf("Cannot open dir '%s' : %w", dirPath, err)
	}
	uploadedFiles := map[string]string{}
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		dirPathArticle := filepath.Join(dirPath, file.Name())
		if err := getUploadArticleFiles(dirPathArticle, &uploadedFiles); err != nil {
			return xerrors.Errorf("Cannot getUploadArticleFiles : %w", err)
		}
	}
	uploadedFiles[fmt.Sprintf("%s/index.json", dirPath)] = "index.json"
	for filePathSrc, filePathDst := range uploadedFiles {
		bytesSrc, err := ioutil.ReadFile(filePathSrc)
		if err != nil {
			return xerrors.Errorf("Cannot read file '%s' : %w", filePathSrc, err)
		}
		if err := articleStore.PutRawFile(ctx, bytesSrc, filePathDst); err != nil {
			return xerrors.Errorf("Cannot put raw file %s => %s : %w", filePathSrc, filePathDst, err)
		}
		fmt.Printf("%s => %s\n", filePathSrc, filePathDst)
	}
	return nil
}

func getUploadArticleFiles(
	dirPathArticle string,
	r *map[string]string,
) error {
	filePathArticleJSON := filepath.Join(dirPathArticle, "article.json")
	article := entity.Article{}
	if err := newArticleFromFile(filePathArticleJSON, &article); err != nil {
		return xerrors.Errorf("Cannot newArticleFromFile : %w", err)
	}
	fileInfos, err := ioutil.ReadDir(dirPathArticle)
	if err != nil {
		return xerrors.Errorf("Cannot ReadDir '%s' : %w", dirPathArticle, err)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		filePathSrc := filepath.Join(dirPathArticle, fileInfo.Name())
		filePathDst := filepath.Join(string(article.ID), fileInfo.Name())
		(*r)[filePathSrc] = filePathDst
	}
	for _, block := range article.Blocks {
		dirPathArticleBlock := filepath.Join(dirPathArticle, string(block.ID))
		if err := getUploadArticleBlockFiles(dirPathArticleBlock, article.ID, block.ID, r); err != nil {
			return xerrors.Errorf("Cannot getUploadArticleBlockFiles : %w", err)
		}
	}
	return nil
}

func getUploadArticleBlockFiles(
	dirPathArticleBlock string,
	articleID entity.ArticleID,
	articleBlockID entity.ArticleBlockID,
	r *map[string]string,
) error {
	fileInfos, err := ioutil.ReadDir(dirPathArticleBlock)
	if err != nil {
		return xerrors.Errorf("Cannot ReadDir '%s' : %w", dirPathArticleBlock, err)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		filePathSrc := filepath.Join(dirPathArticleBlock, fileInfo.Name())
		filePathDst := filepath.Join(string(articleID), string(articleBlockID), fileInfo.Name())
		(*r)[filePathSrc] = filePathDst
	}
	return nil
}
