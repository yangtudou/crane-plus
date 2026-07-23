package planner

import (
	"fmt"
	"strings"

	"registry-sync/model"
)

func Dump(plans []model.Plan) string {

	var b strings.Builder

	b.WriteString("PLAN\n")
	b.WriteString("====\n\n")

	if len(plans) == 0 {
		return b.String()
	}

	// mirrors 只输出一次
	b.WriteString("MIRRORS:\n")

	for _, mirror := range plans[0].Mirrors {

		if mirror.Type != "" {
			b.WriteString(fmt.Sprintf(
				"  %s (%s)\n",
				mirror.URL,
				mirror.Type,
			))
		} else {
			b.WriteString(fmt.Sprintf(
				"  %s\n",
				mirror.URL,
			))
		}
	}

	b.WriteString("\nIMAGES:\n\n")

	currentRegistry := ""

	for _, plan := range plans {

		if plan.Image.Registry != currentRegistry {
			currentRegistry = plan.Image.Registry

			b.WriteString(currentRegistry)
			b.WriteString("\n")
		}

		b.WriteString("  ")
		b.WriteString(plan.Image.Repository)
		b.WriteString(":")
		b.WriteString(plan.Image.Tag)
		b.WriteString("\n")

		// platform
		if len(plan.Image.Platform) > 0 {

			b.WriteString("    platform:\n")

			for _, platform := range plan.Image.Platform {

				b.WriteString("      - ")
				b.WriteString(platform)
				b.WriteString("\n")
			}
		}

		// targets
		for _, target := range plan.Targets {

			b.WriteString("    -> ")

			b.WriteString(
				buildTargetName(
					plan.Image,
					target,
				),
			)

			b.WriteString("\n")
		}

		b.WriteString("\n")
	}

	return b.String()
}

func buildTargetName(
	image model.Image,
	target model.Target,
) string {

	repository := image.Repository

	if target.Flatten {

		parts := strings.Split(
			repository,
			"/",
		)

		repository = parts[len(parts)-1]
	}

	if target.Namespace != "" {

		repository =
			target.Namespace +
				"/" +
				repository
	}

	result := target.Registry + "/" + repository

	if image.Tag != "" {
		result += ":" + image.Tag
	}

	return result
}
