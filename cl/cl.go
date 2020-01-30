package cl

import (
	"flag"
)

var (
	Args []string
	Put bool
)

func init() {
	flag.BoolVar(&Put, "p", false, "Actually put the file into the content-addressable storage.")

	flag.Parse()

	Args = flag.Args()
}
