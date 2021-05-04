package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/subcommands"
	"github.com/suzuito/s2-demo-go/entity"
	"github.com/suzuito/s2-demo-go/usecase"
)

type buildIndexCmd struct {
	dirPathArticles string
}

func newBuildIndexCmd() *buildIndexCmd {
	return &buildIndexCmd{
		dirPathArticles: "",
	}
}

func (c *buildIndexCmd) Name() string     { return "build-index" }
func (c *buildIndexCmd) Synopsis() string { return "Build-index subcomment." }
func (c *buildIndexCmd) Usage() string {
	return "Build-index subcomment.\n"
}

func (c *buildIndexCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.dirPathArticles, "input-dir", "", "Base directory")
}

func (c *buildIndexCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.dirPathArticles == "" {
		fmt.Fprintf(os.Stderr, "dirPathArticles is empty\n")
		return subcommands.ExitUsageError
	}
	root := entity.ArticleListItem{}
	if err := usecase.BuildIndexToLocal(c.dirPathArticles, &root); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return subcommands.ExitFailure
	}
	out, err := json.MarshalIndent(usecase.NewResponseArticleListItem(&root), "", " ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return subcommands.ExitFailure
	}
	if err := ioutil.WriteFile(
		filepath.Join(c.dirPathArticles, "index.json"),
		out,
		0644,
	); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
