package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/suzuito/s2-demo-go/usecase"
)

type buildCmd struct {
	dirPathArticle string
}

func newBuildCmd() *buildCmd {
	return &buildCmd{
		dirPathArticle: "",
	}
}

func (c *buildCmd) Name() string     { return "build" }
func (c *buildCmd) Synopsis() string { return "Build subcomment." }
func (c *buildCmd) Usage() string {
	return "Build subcomment.\n"
}

func (c *buildCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.dirPathArticle, "input-article-dir", "", "Base directory of articles")
}

func (c *buildCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.dirPathArticle == "" {
		fmt.Fprintf(os.Stderr, "dirPathArticle is empty\n")
		return subcommands.ExitUsageError
	}
	if err := usecase.BuildArticleToLocal(c.dirPathArticle); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
