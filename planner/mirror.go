package planner

import (
	"registry-sync/config"
	"registry-sync/model"
)

func resolveMirrors(cfg *config.Config) []model.Mirror {

	if len(cfg.Mirror) == 0 {
		return nil
	}

	mirrors := make([]model.Mirror, 0, len(cfg.Mirror))

	for _, item := range cfg.Mirror {

		mirrors = append(mirrors, model.Mirror{
			Name: item.Name,
			URL:  item.URL,
			Type: model.MirrorType(item.Type),
			Auth: convertAuth(item.Auth),
		})
	}

	return mirrors
}
