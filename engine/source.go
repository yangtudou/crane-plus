package engine

import (
	"registry-sync/model"
)

func ResolveSource(plan model.Plan) string {

	return BuildImageName(
		plan.Image.Registry,
		plan.Image.Repository,
		plan.Image.Tag,
	)
}
