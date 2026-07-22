package main

import (
	"fmt"
	"log"

	"github.com/yyysay/registry-sync/internal/config"
	"github.com/yyysay/registry-sync/internal/destination"
	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/mapper"
	"github.com/yyysay/registry-sync/internal/task"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	images, err := image.Load("images.txt")
	if err != nil {
		log.Fatal(err)
	}

	dst := destination.New(
		cfg.Destination.Registry,
		mapper.New(cfg.Destination.RepositoryMode()),
	)

	tasks := task.Generate(images, dst)

	for _, t := range tasks {
		fmt.Printf(
			"%s -> %s\n",
			t.Source.Reference,
			t.Target.Reference,
		)
	}
}
