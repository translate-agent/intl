package main

import (
	"flag"
)

func main() {
	cldrDir := flag.String("cldr-dir", "", "path to CLDR directory")
	out := flag.String("out", "", "output directory")

	flag.Parse()

	if *cldrDir == "" {
		panic("-cldr-dir flag is required")
	}

	if *out == "" {
		panic("-out flag is required")
	}

	if err := Gen(*cldrDir, *out); err != nil {
		panic(err)
	}
}
