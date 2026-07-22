package mapper

import (
	"path"

	"github.com/yyysay/registry-sync/internal/image"
)

type RepositoryMode int

const (
	Preserve RepositoryMode = iota
	Basename
)

type Mapper struct {
	Mode RepositoryMode
}

func New(mode RepositoryMode) *Mapper {
	return &Mapper{
		Mode: mode,
	}
}

func (m *Mapper) Map(src *image.Image) *image.Image {
	target := *src

	switch m.Mode {
	case Basename:
		target.Name = path.Base(src.Name)
	}

	target.Reference = buildReference(&target)

	return &target
}

func buildReference(img *image.Image) string {
	ref := ""

	if img.Registry != "" {
		ref += img.Registry + "/"
	}

	ref += img.Name

	if img.Tag != "" {
		ref += ":" + img.Tag
	}

	if img.Digest != "" {
		ref += "@" + img.Digest
	}

	return ref
}
