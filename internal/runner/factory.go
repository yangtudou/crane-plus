package runner

import "fmt"

func New(
	name string,
	platform string,
	workers int,
) (Runner, error) {

	switch name {

	case "crane":

		return NewCrane(platform, workers), nil

	case "image-syncer":

		return NewImageSyncer(), nil

	default:

		return nil, fmt.Errorf(
			"unsupported runner: %s",
			name,
		)
	}
}
