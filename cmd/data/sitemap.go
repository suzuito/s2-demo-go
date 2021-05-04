package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/suzuito/s2-demo-go/entity"
	"github.com/suzuito/s2-demo-go/inject"
	"github.com/suzuito/s2-demo-go/sgcp"
	"github.com/suzuito/s2-demo-go/usecase"
)

type sitemapCmd struct {
	dirPathArticles string
}

func newSitemapCmd() *sitemapCmd {
	return &sitemapCmd{
		dirPathArticles: "",
	}
}

func (c *sitemapCmd) Name() string     { return "sitemap" }
func (c *sitemapCmd) Synopsis() string { return "Sitemap subcomment." }
func (c *sitemapCmd) Usage() string {
	return "Sitemap subcomment.\n"
}

func (c *sitemapCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.dirPathArticles, "input-dir", "", "Base directory of articles")
}

func (c *sitemapCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.dirPathArticles == "" {
		fmt.Fprintf(os.Stderr, "dirPathArticles is empty\n")
		return subcommands.ExitUsageError
	}
	env, genGCP, err := inject.NewImplement()
	if err != nil {
		fmt.Printf("NewImplement is failed : %+v\n", err)
		return subcommands.ExitFailure
	}
	resource, err := genGCP.Gen(ctx)
	if err != nil {
		fmt.Printf("Gen resource is failed : %+v\n", err)
		return subcommands.ExitFailure
	}
	store := sgcp.NewServerStore(resource.GCS, env.GCPBucketServer)
	root := entity.ArticleListItem{}
	if err := usecase.BuildIndexToLocal(c.dirPathArticles, &root); err != nil {
		fmt.Printf("Cannot build index : %+v\n", err)
		return subcommands.ExitFailure
	}
	x := entity.XMLURLSet{}
	if err := usecase.GenerateBlogSiteMap(
		ctx,
		&root,
		env.OriginServer,
		&x,
	); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return subcommands.ExitFailure
	}
	if err := store.PutSitemap(ctx, env.OriginServer, &x); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
