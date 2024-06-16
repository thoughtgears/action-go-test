package main

import (
	"context"

	"github.com/sethvargo/go-githubactions"
	"github.com/thoughtgears/iac-generator/pkg/generator"
)

func run() error {
	ctx := context.Background()
	action := githubactions.New()

	cfg, err := generator.NewFromInputs(action)
	if err != nil {
		return err
	}

	return generator.Run(ctx, cfg)
}

func main() {
	err := run()
	if err != nil {
		githubactions.Fatalf("%v", err)
	}
}
