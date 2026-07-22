package main

import (
	"github.com/yyysay/registry-sync/internal/config"
	"github.com/yyysay/registry-sync/internal/destination"
	"github.com/yyysay/registry-sync/internal/mapper"
	"github.com/yyysay/registry-sync/internal/rule"
	"github.com/yyysay/registry-sync/internal/source"
)

func buildRules(
	cfg *config.Config,
) ([]*rule.Rule, error) {

	var rules []*rule.Rule

	for _, item := range cfg.Rules {

		mode := mapper.Basename

		if item.Destination.Mode == "preserve" {

			mode = mapper.Preserve

		}

		var src *source.Source

		// source.registry 存在:
		// 普通规则
		//
		// source.registry 为空:
		// mixed fallback

		if item.Source.Registry != "" {

			src = source.New(
				item.Source.Registry,
				source.Default,
			)

		}

		rules = append(
			rules,

			rule.New(
				item.Name,

				src,

				destination.New(
					item.Destination.Registry,
					mapper.New(mode),
				),
			),
		)

	}

	return rules, nil
}
