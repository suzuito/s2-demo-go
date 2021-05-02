package sgcp

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/suzuito/s2-demo-go/entity"
	"golang.org/x/xerrors"
	"google.golang.org/api/iterator"
)

type ArticleStore struct {
	cli    *storage.Client
	bucket string
}

func NewArticleStore(cli *storage.Client, bucket string) *ArticleStore {
	return &ArticleStore{
		cli:    cli,
		bucket: bucket,
	}
}

func (s *ArticleStore) GetArticle(
	ctx context.Context,
	articleID entity.ArticleID,
	article *entity.Article,
) error {
	bh := s.cli.Bucket(s.bucket)
	keyBase := filepath.Join(string(articleID))
	it := bh.Objects(ctx, &storage.Query{
		Prefix: keyBase,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return xerrors.Errorf("%s/%s : %w", s.bucket, keyBase, err)
		}
		block := entity.ArticleBlock{}
		if err := s.getArticle(ctx, bh, articleID, attrs.Name, &block); err != nil {
			return xerrors.Errorf("Cannot get article : %w", err)
		}
	}
	return nil
}

func (s *ArticleStore) getArticle(
	ctx context.Context,
	bh *storage.BucketHandle,
	articleID entity.ArticleID,
	blockID string,
	block *entity.ArticleBlock,
) error {
	keyBase := filepath.Join(string(articleID), blockID)
	it := bh.Objects(ctx, &storage.Query{
		Prefix: keyBase,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return xerrors.Errorf("%s : %w", s.bucket, attrs.Name, err)
		}
		// key := filepath.Join(attrs.Prefix, attrs.Name)
		oh := bh.Object(attrs.Name)
		reader, err := oh.NewReader(ctx)
		if err != nil {
			return xerrors.Errorf("Cannot NewReader '%s' : %w", attrs.Name, err)
		}
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			return xerrors.Errorf("Cannot ReadAll '%s' : %w", attrs.Name, err)
		}
		fmt.Printf("%s\n%s\n\n", attrs.Name, string(body))
	}
	return nil
}

func (s *ArticleStore) GetArticleList(
	ctx context.Context,
	indecies *entity.ArticleListItem,
) error {
	return xerrors.Errorf("Not impl")
}

func (s *ArticleStore) PutArticle(
	ctx context.Context,
	article *entity.Article,
) error {
	bh := s.cli.Bucket(s.bucket)
	keyArticleDir := filepath.Join(string(article.ID))
	for _, block := range article.Blocks {
		if err := putArticleBlock(ctx, bh, keyArticleDir, &block); err != nil {
			return xerrors.Errorf("Cannot put article block : %w", err)
		}
	}
	return nil
}

func putArticleBlock(
	ctx context.Context,
	bh *storage.BucketHandle,
	keyArticleDir string,
	block *entity.ArticleBlock,
) error {
	if block.Text != "" {
		k := filepath.Join(keyArticleDir, "article.html")
		w := bh.Object(k).NewWriter(ctx)
		defer w.Close()
		if _, err := fmt.Fprintf(w, block.Text); err != nil {
			return xerrors.Errorf("Cannot write '%s' : %w", k, err)
		}
	}
	if block.Source != "" {
		k := filepath.Join(keyArticleDir, "main.go")
		w := bh.Object(k).NewWriter(ctx)
		defer w.Close()
		if _, err := fmt.Fprintf(w, block.Source); err != nil {
			return xerrors.Errorf("Cannot write '%s' : %w", k, err)
		}
	}
	return nil
}
