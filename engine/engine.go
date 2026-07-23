package engine

import (
	"context"
	"fmt"

	"registry-sync/model"
)

type Engine struct {
	copier Copier
}

func New(copier Copier) *Engine {
	return &Engine{
		copier: copier,
	}
}

func (e *Engine) Execute(
	ctx context.Context,
	plan model.Plan,
) error {

	source := ResolveSource(plan)

	fmt.Println("SOURCE:")
	fmt.Println(" ", source)

	for _, target := range plan.Targets {

		targetImage := BuildTargetImage(
			plan.Image,
			target,
		)

		fmt.Println("COPY:")
		fmt.Println(" ", source)
		fmt.Println(" ->")
		fmt.Println(" ", targetImage)

		err := e.copier.Copy(
			ctx,
			source,
			targetImage,
			plan.Image.Platform,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
