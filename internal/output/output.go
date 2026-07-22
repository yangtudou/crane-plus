package output

import (
	"fmt"

	"github.com/yyysay/registry-sync/internal/task"
)

func Print(tasks []*task.Task) {
	for _, t := range tasks {
		fmt.Printf(
			"%s -> %s\n",
			t.Source.Reference,
			t.Target.Reference,
		)
	}
}
