package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yyysay/registry-sync/internal/config"
	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/validate"
)

func validateCommand(args []string) {
	fs := flag.NewFlagSet(
		"validate",
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

	if err := validate.Config(cfg); err != nil {
		log.Fatal(err)
	}

	images, err := image.Load(*imageFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := validate.Images(images); err != nil {
		log.Fatal(err)
	}

	fmt.Println("config ok")
	fmt.Println("images ok")
}
