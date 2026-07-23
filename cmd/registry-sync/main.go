package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"registry-sync/config"
	"registry-sync/copier"
	"registry-sync/engine"
	"registry-sync/planner"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("usage: registry-sync <config.yaml>")
	}

	configFile := os.Args[1]

	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	plans := planner.Build(cfg)

	fmt.Println(
		planner.Dump(plans),
	)

	e := engine.New(
		copier.New(),
	)

	ctx := context.Background()

	for _, plan := range plans {

		err := e.Execute(
			ctx,
			plan,
		)

		if err != nil {
			log.Fatalf(
				"sync failed: %v",
				err,
			)
		}
	}
}
