package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yyysay/registry-sync/internal/config"
	"github.com/yyysay/registry-sync/internal/destination"
	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/mapper"
	"github.com/yyysay/registry-sync/internal/output"
	"github.com/yyysay/registry-sync/internal/task"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {
	case "plan":
		plan(os.Args[2:])
	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}
}

func usage() {
	fmt.Println("usage:")
	fmt.Println("  registry-sync plan [options]")
	fmt.Println()
	fmt.Println("commands:")
	fmt.Println("  plan")
}

func plan(args []string) {
	fs := flag.NewFlagSet("plan", flag.ExitOnError)

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

	dst := destination.New(
		cfg.Destination.Registry,
		mapper.New(cfg.Destination.RepositoryMode()),
	)

	tasks := task.Generate(images, dst)

	output.Print(tasks)
}
