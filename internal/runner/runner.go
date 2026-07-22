package runner

import (
	"github.com/yyysay/registry-sync/internal/task"
)

type Runner interface {
	Run(tasks []*task.Task) error
}
