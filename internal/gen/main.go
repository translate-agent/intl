package main

import (
	"flag"
)

func main() {
	dir := flag.String("dir", "", "path to CLDR")

	flag.Parse()

	if *dir == "" {
		panic("-dir flag is required")
	}

	if err := Gen(*dir); err != nil {
		panic(err)
	}
}
