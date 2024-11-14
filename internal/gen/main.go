package main

import (
	"flag"
	"fmt"
)

func main() {
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

	if err := Gen(conf); err != nil {
		panic(err)
	}

	fmt.Println("done")
}
