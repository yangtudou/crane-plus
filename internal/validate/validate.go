package validate

import (
	"fmt"

	"github.com/yyysay/registry-sync/internal/config"
	"github.com/yyysay/registry-sync/internal/image"
)

func Config(cfg *config.Config) error {
	if cfg.Destination.Registry == "" {
		return fmt.Errorf("destination.registry is empty")
	}

	switch cfg.Destination.Mode {
	case "basename", "preserve":
	default:
		return fmt.Errorf(
			"unsupported destination.mode: %s",
			cfg.Destination.Mode,
		)
	}

	return nil
}

func Images(images []*image.Image) error {
	if len(images) == 0 {
		return fmt.Errorf("image list is empty")
	}

	return nil
}
