package main

import (
	"flag"
	"log"

	"github.com/yyysay/registry-sync/internal/config"
	"github.com/yyysay/registry-sync/internal/engine"
	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/output"
)

func plan(args []string) {

	fs := flag.NewFlagSet(
		"plan",
		flag.ExitOnError,
	)

	configFile := fs.String(
		"config",
		"config.yaml",
		"config file",
	)

	imageFile := fs.String(
		"images",
		"images.txt",
		"image list file",
	)

	format := fs.String(
		"format",
		"text",
		"output format: text|json",
	)

	fs.Parse(args)

	cfg, err := config.Load(*configFile)

	if err != nil {
		log.Fatal(err)
	}

	images, err := image.Load(*imageFile)

	if err != nil {
		log.Fatal(err)
	}

	rules, err := buildRules(cfg)

	if err != nil {
		log.Fatal(err)
	}

	tasks, err := engine.
		New(rules).
		Generate(images)

	if err != nil {
		log.Fatal(err)
	}

	switch *format {

	case "json":

		if err := output.PrintJSON(tasks); err != nil {
			log.Fatal(err)
		}

	default:

		output.PrintText(tasks)

	}
}
