package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
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
	if err := usecase.BuildIndexToLocal(c.dirPathArticles); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
