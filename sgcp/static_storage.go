package sgcp

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/suzuito/s2-demo-go/entity"
	"golang.org/x/xerrors"
)

type ServerStore struct {
	cli    *storage.Client
	bucket string
}

func NewServerStore(cli *storage.Client, bucket string) *ServerStore {
	return &ServerStore{
		cli:    cli,
		bucket: bucket,
	}
}

func (s *ServerStore) PutSitemap(
	ctx context.Context,
	origin string,
	x *entity.XMLURLSet,
) error {
	oh := s.cli.Bucket(s.bucket).Object("sitemap.xml")
	w := oh.NewWriter(ctx)
	defer w.Close()
	w.ContentType = "application/xml; charset=utf-8"
	b, err := x.Marshal()
	if err != nil {
		return xerrors.Errorf("Cannot marshal xml : %w", err)
	}
	if _, err := fmt.Fprint(w, b); err != nil {
		return xerrors.Errorf("Cannot write xml : %w", err)
	}
	return nil
}
