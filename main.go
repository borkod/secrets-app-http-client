package main

import (
	"os"

	"github.com/borkod/secrets-http-client/cmd"
)

func main() {
	if err := cmd.Root(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
