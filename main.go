package main

import (
	"os"

	"github.com/matthew-healy/gopack/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
