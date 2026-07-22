package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yyysay/registry-sync/internal/task"
)

type Item struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

func Print(tasks []*task.Task) {
	PrintText(tasks)
}

func PrintText(tasks []*task.Task) {
	for _, t := range tasks {
		fmt.Printf(
			"%s -> %s\n",
			t.Source.Reference,
			t.Target.Reference,
		)
	}
}

func PrintJSON(tasks []*task.Task) error {
	items := make([]Item, 0, len(tasks))

	for _, t := range tasks {
		items = append(items, Item{
			Source: t.Source.Reference,
			Target: t.Target.Reference,
		})
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	return encoder.Encode(items)
}
