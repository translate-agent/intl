package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()

	var conf Conf

	flag.StringVar(&conf.cldrDir, "cldr-dir", "", "path to CLDR directory")
	flag.StringVar(&conf.out, "out", "", "output directory")
	flag.BoolVar(&conf.saveMerged, "save-merged", false, "save merged CLDR data")

	flag.Parse()

	if conf.cldrDir == "" {
		panic("-cldr-dir flag is required")
	}

	if conf.out == "" {
		panic("-out flag is required")
	}

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	err := Gen(ctx, conf, log)
	if err != nil {
		panic(err)
	}

	log.InfoContext(ctx, "done")
}
