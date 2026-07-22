package rule

import (
	"github.com/yyysay/registry-sync/internal/destination"
	"github.com/yyysay/registry-sync/internal/source"
)

type Rule struct {
	Name        string
	Source      *source.Source
	Destination *destination.Destination
}

func New(
	name string,
	src *source.Source,
	dst *destination.Destination,
) *Rule {
	return &Rule{
		Name:        name,
		Source:      src,
		Destination: dst,
	}
}
