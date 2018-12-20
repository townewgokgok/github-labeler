package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/townewgokgok/github-labeler"
)

func main() {
	var (
		manifest = flag.String("manifest", "labels.yaml", "YAML file to be described about labels and repos")
		dryRun   = flag.Bool("dry-run", false, "dry run flag")
	)
	flag.Parse()

	if err := labeler.Run(*manifest, *dryRun); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err.Error())
		os.Exit(1)
	}
}
