package cl

import (
	"os"
)

var (
	CommandName = "tgfs-hash"
)

func init() {
	if 1 > len(os.Args) {
		return
	}

	CommandName = os.Args[0]
}
