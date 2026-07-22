package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {

	case "plan":
		plan(os.Args[2:])

	case "sync":
		syncCommand(os.Args[2:])

	case "validate":
		validateCommand(os.Args[2:])

	default:
		log.Fatalf(
			"unknown command: %s",
			os.Args[1],
		)
	}
}

func usage() {

	fmt.Println("usage:")
	fmt.Println(
		"  registry-sync <command>",
	)

	fmt.Println()

	fmt.Println("commands:")

	fmt.Println(
		"  plan",
	)

	fmt.Println(
		"  sync",
	)

	fmt.Println(
		"  validate",
	)
}
