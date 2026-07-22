package task

import (
	"github.com/yyysay/registry-sync/internal/destination"
	"github.com/yyysay/registry-sync/internal/image"
)

type Task struct {
	Source *image.Image
	Target *image.Image
}

func Generate(images []*image.Image, dst *destination.Destination) []*Task {
	tasks := make([]*Task, 0, len(images))

	for _, img := range images {
		tasks = append(tasks, &Task{
			Source: img,
			Target: dst.Map(img),
		})
	}

	return tasks
}
