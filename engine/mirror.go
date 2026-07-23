package engine

import (
	"registry-sync/model"
)

func BuildMirrorImage(
	image model.Image,
	mirror model.Mirror,
) string {

	repository := image.Repository

	if mirror.Type == "flatten" {
		repository = flattenRepository(repository)
	}

	return BuildImageName(
		mirror.URL,
		repository,
		image.Tag,
	)
}

func BuildMirrorImages(
	image model.Image,
	mirrors []model.Mirror,
) []string {

	var result []string

	for _, mirror := range mirrors {
		result = append(
			result,
			BuildMirrorImage(image, mirror),
		)
	}

	return result
}
