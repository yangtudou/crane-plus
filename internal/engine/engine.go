package engine

import (
	"github.com/yyysay/registry-sync/internal/image"
	"github.com/yyysay/registry-sync/internal/rule"
	"github.com/yyysay/registry-sync/internal/task"
)

type Engine struct {
	rules []*rule.Rule
}

func New(
	rules []*rule.Rule,
) *Engine {

	return &Engine{
		rules: rules,
	}

}

func (e *Engine) Generate(
	images []*image.Image,
) ([]*task.Task, error) {

	var tasks []*task.Task

	for _, img := range images {

		var matched *rule.Rule

		var mixed *rule.Rule

		// 查找规则

		for _, r := range e.rules {

			if r == nil ||
				r.Destination == nil {

				continue

			}

			// 保存 mixed

			if r.IsMixed() {

				mixed = r
				continue

			}

			// 普通规则匹配

			if r.Match(img) {

				matched = r
				break

			}

		}

		// fallback

		if matched == nil {

			matched = mixed

		}

		// 没有任何规则

		if matched == nil {

			continue

		}

		target := matched.Destination.Map(img)

		tasks = append(
			tasks,
			&task.Task{
				Source: img,
				Target: target,
			},
		)

	}

	return tasks, nil
}
