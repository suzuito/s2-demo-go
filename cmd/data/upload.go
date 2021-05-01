package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/google/subcommands"
	"github.com/suzuito/s2-demo-go/usecase"
)

type uploadCmd struct {
	dirPathArticle string
}

func newUploadCmd() *uploadCmd {
	return &uploadCmd{
		dirPathArticle: "",
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
		var err error
		var inputFile *os.File
		var outputFile *os.File
		if info.IsDir() {
			return nil
		}
		fmt.Println(p)
		if filepath.Ext(p) == ".md" {
			err = usecase.ConvertFileMarkdownToHTML(
				p,
				regexp.Compile(),
			)
		}
		// outputFile, err := os.Open()
		// usecase.ConvertMarkdownToHTML(inputFile, )
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot walk : %+v\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
