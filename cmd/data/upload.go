package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/subcommands"
)

type uploadCmd struct {
	dirPathArticle string
}

func newUploadCmd(
	dirPathArticle string,
) *uploadCmd {
	return &uploadCmd{
		dirPathArticle: dirPathArticle,
	}
}

func (c *uploadCmd) Name() string     { return "upload" }
func (c *uploadCmd) Synopsis() string { return "Upload subcomment." }
func (c *uploadCmd) Usage() string {
	return "Upload subcomment.\n"
}

func (c *uploadCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.dirPathArticle, "input-article-dir", "", "Base directory of articles")
}

func (c *uploadCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	err := filepath.Walk(c.dirPathArticle, func(p string, info os.FileInfo, err error) error {
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot walk : %+v\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
