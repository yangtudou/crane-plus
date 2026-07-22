package validate

import (
	"fmt"

	"github.com/yyysay/registry-sync/internal/config"
	"github.com/yyysay/registry-sync/internal/image"
)

func Config(cfg *config.Config) error {
	if len(cfg.Rules) == 0 {
		return fmt.Errorf("rules is empty")
	}

	for _, r := range cfg.Rules {
		if r.Name == "" {
			return fmt.Errorf("rule name is empty")
		}

		if r.Source.Registry == "" {
			return fmt.Errorf(
				"rule %s source.registry is empty",
				r.Name,
			)
		}

		if r.Destination.Registry == "" {
			return fmt.Errorf(
				"rule %s destination.registry is empty",
				r.Name,
			)
		}

		switch r.Destination.Mode {
		case "basename", "preserve":
		default:
			return fmt.Errorf(
				"rule %s unsupported destination.mode: %s",
				r.Name,
				r.Destination.Mode,
			)
		}
	}

	return nil
}

func Images(images []*image.Image) error {
	if len(images) == 0 {
		return fmt.Errorf("image list is empty")
	}

	return nil
}
