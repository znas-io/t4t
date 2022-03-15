package main

import (
	"github.com/znas-io/t4t/cmd/root"
	"os"
)

func main() {
	if err := root.NewCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
