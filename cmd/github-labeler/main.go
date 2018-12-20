package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/townewgokgok/github-labeler"
)

func main() {
	var (
		baseURL  = flag.String("base-url", "", "base URL of Github API")
		token    = flag.String("token", "", "Github access token")
		manifest = flag.String("manifest", "labels.yaml", "YAML file to be described about labels and repos")
		dryRun   = flag.Bool("dry-run", false, "dry run flag")
	)
	flag.Parse()

	opts := labeler.Options{
		BaseURL:    *baseURL,
		Token:      *token,
		ConfigPath: *manifest,
		DryRun:     *dryRun,
	}
	if err := labeler.Run(opts); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err.Error())
		os.Exit(1)
	}
}
