package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/suzuito/s2-demo-go/inject"
	"github.com/suzuito/s2-demo-go/sgcp"
	"github.com/suzuito/s2-demo-go/usecase"
)

type uploadCmd struct {
	dirPathArticles string
}

func newUploadCmd() *uploadCmd {
	return &uploadCmd{
		dirPathArticles: "",
	}
}

func (c *uploadCmd) Name() string     { return "upload" }
func (c *uploadCmd) Synopsis() string { return "Upload subcomment." }
func (c *uploadCmd) Usage() string {
	return "Upload subcomment.\n"
}

func (c *uploadCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.dirPathArticles, "input-dir", "", "Base directory of articles")
}

func (c *uploadCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
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
	articleStore := sgcp.NewArticleStore(resource.GCS, env.GCPBucketArticle)
	if err := usecase.UploadArticles(
		ctx,
		articleStore,
		c.dirPathArticles,
	); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
