package source

import "github.com/yyysay/registry-sync/internal/image"

type Mode int

const (
	Default Mode = iota
)

type Source struct {
	Registry string
	Mode     Mode
}

func New(
	registry string,
	mode Mode,
) *Source {
	return &Source{
		Registry: registry,
		Mode:     mode,
	}
}

func (s *Source) Match(img *image.Image) bool {
	if s == nil || img == nil {
		return false
	}

	return img.Registry == s.Registry
}
