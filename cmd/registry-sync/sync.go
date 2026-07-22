package main

import (
	"flag"
	"log"

	"github.com/yyysay/registry-sync/internal/config"
	"github.com/yyysay/registry-sync/internal/engine"
	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/runner"
)

func syncCommand(args []string) {

	fs := flag.NewFlagSet(
		"sync",
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

	runnerType := cfg.Runner.Type

	// 兼容旧配置
	if runnerType == "" {
		runnerType = "crane"
	}

	r, err := runner.New(
		runnerType,
		cfg.Runner.Platform,
		cfg.Runner.Workers,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := r.Run(tasks); err != nil {
		log.Fatal(err)
	}
}
