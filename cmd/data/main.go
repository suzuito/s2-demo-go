package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

// uploader は、GeoJSONファイルを PostGIS へ挿れるためのコマンド
func main() {
	ctx := context.Background()
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(newBuildCmd(), "")
	subcommands.Register(newBuildIndexCmd(), "")
	subcommands.Register(newUploadCmd(), "")
	subcommands.Register(newSitemapCmd(), "")
	flag.Parse()

	os.Exit(int(subcommands.Execute(ctx)))
}
