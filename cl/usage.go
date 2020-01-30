package cl

import (
	"github.com/reiver/tgfs-hash/hash"

	"flag"
	"fmt"
)

func Usage() {
	w := flag.CommandLine.Output()

	fmt.Fprintf(w, "Compute the %s hash function digest for a file, and optionally add the file to a local instance of The Great File System (TGFS).\n\n", hash.Name)
	fmt.Fprint(w, "Examples:\n")
	fmt.Fprintf(w, "\t%s <filename>\n\n", CommandName)
	fmt.Fprintf(w, "\tcat <filename> | %s\n\n", CommandName)
	fmt.Fprintf(w, "\t%s -w <filename>\n\n", CommandName)
	fmt.Fprintf(w, "\tcat <filename> | %s -w\n\n", CommandName)
	flag.Usage()
}
