package main

import (
	"testing"

	"github.com/yyysay/registry-sync/internal/config"
)

func TestBuildRulesMixed(t *testing.T) {

	cfg := &config.Config{

		Rules: []config.Rule{

			{
				Name: "mixed",

				Destination: config.Destination{
					Registry: "registry.example.com/images",
					Mode:     "basename",
				},
			},
		},
	}

	rules, err := buildRules(cfg)

	if err != nil {
		t.Fatal(err)
	}

	if len(rules) != 1 {

		t.Fatalf(
			"rules=%d",
			len(rules),
		)

	}

	if !rules[0].IsMixed() {

		t.Fatal(
			"expected mixed rule",
		)

	}

}
